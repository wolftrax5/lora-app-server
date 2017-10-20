package storage

import (
	"context"
	"time"

	"github.com/brocaar/loraserver/api/ns"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/brocaar/lora-app-server/internal/common"
	"github.com/brocaar/lorawan"
)

// Device defines a LoRaWAN device.
type Device struct {
	DevEUI          lorawan.EUI64 `db:"dev_eui"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
	ApplicationID   int64         `db:"application_id"`
	DeviceProfileID string        `db:"device_profile_id"`
	Name            string        `db:"name"`
	Description     string        `db:"description"`
}

// Validate validates the device data.
func (d Device) Validate() error {
	return nil
}

// DeviceKeys defines the keys for a LoRaWAN device.
type DeviceKeys struct {
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
	DevEUI    lorawan.EUI64     `db:"dev_eui"`
	AppKey    lorawan.AES128Key `db:"app_key"`
	JoinNonce int               `db:"join_nonce"`
}

// DeviceActivation defines the device-activation for a LoRaWAN device.
type DeviceActivation struct {
	ID        int64             `db:"id"`
	CreatedAt time.Time         `db:"created_at"`
	DevEUI    lorawan.EUI64     `db:"dev_eui"`
	DevAddr   lorawan.DevAddr   `db:"dev_addr"`
	AppSKey   lorawan.AES128Key `db:"app_s_key"`
	NwkSKey   lorawan.AES128Key `db:"nwk_s_key"`
}

// CreateDevice creates the given device.
func CreateDevice(db sqlx.Ext, d *Device) error {
	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	now := time.Now()
	d.CreatedAt = now
	d.UpdatedAt = now

	_, err := db.Exec(`
        insert into device (
            dev_eui,
            created_at,
            updated_at,
            application_id,
            device_profile_id,
            name,
            description
        ) values ($1, $2, $3, $4, $5, $6, $7)`,
		d.DevEUI[:],
		d.CreatedAt,
		d.UpdatedAt,
		d.ApplicationID,
		d.DeviceProfileID,
		d.Name,
		d.Description,
	)
	if err != nil {
		return handlePSQLError(err, "insert error")
	}

	app, err := GetApplication(db, d.ApplicationID)
	if err != nil {
		return errors.Wrap(err, "get application error")
	}

	n, err := GetNetworkServerForDevEUI(db, d.DevEUI)
	if err != nil {
		return errors.Wrap(err, "get network-server error")
	}

	nsClient, err := common.NetworkServerPool.Get(n.Server)
	if err != nil {
		return errors.Wrap(err, "get network-server client error")
	}

	_, err = nsClient.CreateDevice(context.Background(), &ns.CreateDeviceRequest{
		Device: &ns.Device{
			DevEUI:           d.DevEUI[:],
			DeviceProfileID:  d.DeviceProfileID,
			ServiceProfileID: app.ServiceProfileID,
			RoutingProfileID: common.ApplicationServerID,
		},
	})
	if err != nil {
		log.WithError(err).Error("network-server create device api error")
		return handleGrpcError(err, "create device error")
	}

	log.WithFields(log.Fields{
		"dev_eui": d.DevEUI,
	}).Info("device created")

	return nil
}

// GetDevice returns the device matching the given DevEUI.
func GetDevice(db sqlx.Queryer, devEUI lorawan.EUI64) (Device, error) {
	var d Device
	err := sqlx.Get(db, &d, "select * from device where dev_eui = $1", devEUI[:])
	if err != nil {
		return d, handlePSQLError(err, "select error")
	}

	return d, nil
}

// GetDevicesForApplicationID returns a slice of devices for the given
// application id.
func GetDevicesForApplicationID(db sqlx.Queryer, applicationID int64, limit, offset int, search string) ([]Device, error) {
	var devices []Device
	if search != "" {
		search = search + "%"
	}
	err := sqlx.Select(db, &devices, `
		select
			*
		from device
		where
			application_id = $1
			and ( ($4 = '') or ($4 != '' and (name ilike $4 or encode(dev_eui, 'hex') ilike $4)) )
		order by name
		limit $2
		offset $3`,
		applicationID,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, handlePSQLError(err, "select error")
	}

	return devices, nil
}

// GetDeviceCountForApplicationID returns the total number of devices for the
// given application id.
func GetDeviceCountForApplicationID(db sqlx.Queryer, applicationID int64, search string) (int, error) {
	var count int
	if search != "" {
		search = search + "%"
	}
	err := sqlx.Get(db, &count, `
		select
			count(*)
		from device
		where
			application_id = $1
			and ( ($2 = '') or ($2 != '' and (name ilike $2 or encode(dev_eui, 'hex') ilike $2)) )`,
		applicationID,
		search,
	)
	if err != nil {
		return count, handlePSQLError(err, "select error")
	}

	return count, nil
}

// UpdateDevice updates the given device.
func UpdateDevice(db sqlx.Ext, d *Device) error {
	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	d.UpdatedAt = time.Now()

	res, err := db.Exec(`
        update device
        set
            updated_at = $2,
            application_id = $3,
            device_profile_id = $4,
            name = $5,
            description = $6
        where
            dev_eui = $1`,
		d.DevEUI[:],
		d.UpdatedAt,
		d.ApplicationID,
		d.DeviceProfileID,
		d.Name,
		d.Description,
	)
	if err != nil {
		return handlePSQLError(err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	app, err := GetApplication(db, d.ApplicationID)
	if err != nil {
		return errors.Wrap(err, "get application error")
	}

	n, err := GetNetworkServerForDevEUI(db, d.DevEUI)
	if err != nil {
		return errors.Wrap(err, "get network-server error")
	}

	nsClient, err := common.NetworkServerPool.Get(n.Server)
	if err != nil {
		return errors.Wrap(err, "get network-server client error")
	}

	_, err = nsClient.UpdateDevice(context.Background(), &ns.UpdateDeviceRequest{
		Device: &ns.Device{
			DevEUI:           d.DevEUI[:],
			DeviceProfileID:  d.DeviceProfileID,
			ServiceProfileID: app.ServiceProfileID,
			RoutingProfileID: common.ApplicationServerID,
		},
	})
	if err != nil {
		log.WithError(err).Error("network-server update device api error")
		return handleGrpcError(err, "update device error")
	}

	log.WithFields(log.Fields{
		"dev_eui": d.DevEUI,
	}).Info("device updated")

	return nil
}

// DeleteDevice deletes the device matching the given DevEUI.
func DeleteDevice(db sqlx.Ext, devEUI lorawan.EUI64) error {
	n, err := GetNetworkServerForDevEUI(db, devEUI)
	if err != nil {
		return errors.Wrap(err, "get network-server error")
	}

	res, err := db.Exec("delete from device where dev_eui = $1", devEUI[:])
	if err != nil {
		return handlePSQLError(err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	nsClient, err := common.NetworkServerPool.Get(n.Server)
	if err != nil {
		return errors.Wrap(err, "get network-server client error")
	}

	_, err = nsClient.DeleteDevice(context.Background(), &ns.DeleteDeviceRequest{
		DevEUI: devEUI[:],
	})
	if err != nil {
		log.WithError(err).Error("network-server delete device api error")
		return handleGrpcError(err, "delete device error")
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
	}).Info("device deleted")

	return nil
}

// CreateDeviceKeys creates the keys for the given device.
func CreateDeviceKeys(db sqlx.Execer, dc *DeviceKeys) error {
	now := time.Now()
	dc.CreatedAt = now
	dc.UpdatedAt = now

	_, err := db.Exec(`
        insert into device_keys (
            created_at,
            updated_at,
            dev_eui,
			app_key,
			join_nonce
        ) values ($1, $2, $3, $4, $5)`,
		dc.CreatedAt,
		dc.UpdatedAt,
		dc.DevEUI[:],
		dc.AppKey[:],
		dc.JoinNonce,
	)
	if err != nil {
		return handlePSQLError(err, "insert error")
	}

	log.WithFields(log.Fields{
		"dev_eui": dc.DevEUI,
	}).Info("device-keys created")

	return nil
}

// GetDeviceKeys returns the device-keys for the given DevEUI.
func GetDeviceKeys(db sqlx.Queryer, devEUI lorawan.EUI64) (DeviceKeys, error) {
	var dc DeviceKeys

	err := sqlx.Get(db, &dc, "select * from device_keys where dev_eui = $1", devEUI[:])
	if err != nil {
		return dc, handlePSQLError(err, "select error")
	}

	return dc, nil
}

// UpdateDeviceKeys updates the given device-keys.
func UpdateDeviceKeys(db sqlx.Execer, dc *DeviceKeys) error {
	dc.UpdatedAt = time.Now()

	res, err := db.Exec(`
        update device_keys
        set
            updated_at = $2,
			app_key = $3,
			join_nonce = $4
        where
            dev_eui = $1`,
		dc.DevEUI[:],
		dc.UpdatedAt,
		dc.AppKey[:],
		dc.JoinNonce,
	)
	if err != nil {
		return handlePSQLError(err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"dev_eui": dc.DevEUI,
	}).Info("device-keys updated")

	return nil
}

// DeleteDeviceKeys deletes the device-keys for the given DevEUI.
func DeleteDeviceKeys(db sqlx.Execer, devEUI lorawan.EUI64) error {
	res, err := db.Exec("delete from device_keys where dev_eui = $1", devEUI[:])
	if err != nil {
		return handlePSQLError(err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected errro")
	}
	if ra == 0 {
		return ErrDoesNotExist
	}

	log.WithField("dev_eui", devEUI).Info("device-keys deleted")

	return nil
}

// CreateDeviceActivation creates the given device-activation.
func CreateDeviceActivation(db sqlx.Queryer, da *DeviceActivation) error {
	da.CreatedAt = time.Now()

	err := sqlx.Get(db, &da.ID, `
        insert into device_activation (
            created_at,
            dev_eui,
            dev_addr,
            app_s_key,
            nwk_s_key
        ) values ($1, $2, $3, $4, $5)
        returning id`,
		da.CreatedAt,
		da.DevEUI[:],
		da.DevAddr[:],
		da.AppSKey[:],
		da.NwkSKey[:],
	)
	if err != nil {
		return handlePSQLError(err, "insert error")
	}

	log.WithFields(log.Fields{
		"id":      da.ID,
		"dev_eui": da.DevEUI,
	}).Info("device-activation created")

	return nil
}

// GetLastDeviceActivationForDevEUI returns the most recent device-activation for the given DevEUI.
func GetLastDeviceActivationForDevEUI(db sqlx.Queryer, devEUI lorawan.EUI64) (DeviceActivation, error) {
	var da DeviceActivation

	err := sqlx.Get(db, &da, `
        select *
        from device_activation
        where
            dev_eui = $1
        order by
            created_at desc
        limit 1`,
		devEUI[:],
	)
	if err != nil {
		return da, handlePSQLError(err, "select error")
	}

	return da, nil
}
