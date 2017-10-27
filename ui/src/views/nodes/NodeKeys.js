import React, { Component } from 'react';

import NodeStore from "../../stores/NodeStore";


class DeviceKeysForm extends Component {
  static contextTypes = {
    router: React.PropTypes.object.isRequired
  };

  constructor() {
    super();

    this.state = {
      deviceKeys: {
        deviceKeys: {},
      },
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  componentDidMount() {
    this.setState({
      deviceKeys: this.props.deviceKeys,
    });
  }

  componentWillReceiveProps(nextProps) {
    this.setState({
      deviceKeys: nextProps.deviceKeys,
    });
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.onSubmit(this.state.deviceKeys);
  }

  onChange(fieldLookup, e) {
    let lookup = fieldLookup.split(".");
    const fieldName = lookup[lookup.length-1];
    lookup.pop(); // remove last item

    let deviceKeys = this.state.deviceKeys;
    let obj = deviceKeys;

    for (const f of lookup) {
      obj = obj[f];
    }

    obj[fieldName] = e.target.value;

    this.setState({
      deviceKeys: deviceKeys,
    });
  }

  render() {
    return(
      <form onSubmit={this.handleSubmit}>
        <div className="form-group">
          <label className="control-label" htmlFor="devEUI">Application key</label>
          <input className="form-control" id="appKey" type="text" placeholder="00000000000000000000000000000000" pattern="[A-Fa-f0-9]{32}" required value={this.state.deviceKeys.deviceKeys.appKey || ''} onChange={this.onChange.bind(this, 'deviceKeys.appKey')} /> 
        </div>
        <hr />
        <div className="btn-toolbar pull-right">
          <a className="btn btn-default" onClick={this.context.router.goBack}>Go back</a>
          <button type="submit" className={"btn btn-primary " + (this.state.disabled ? 'hidden' : '')}>Submit</button>
        </div>
      </form>
    );
  }
}

class NodeKeys extends Component {
  static contextTypes = {
    router: React.PropTypes.object.isRequired
  };

  constructor() {
    super();

    this.state = {
      deviceKeys: {
        deviceKeys: {},
      },
      update: false,
    };

    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    NodeStore.getNodeKeys(this.props.params.devEUI, (deviceKeys) => {
      this.setState({
        update: true,
        deviceKeys: deviceKeys,
      });
    });
  }

  onSubmit(deviceKeys) {
    if (this.state.update) {
      NodeStore.updateNodeKeys(this.props.params.devEUI, deviceKeys, (responseData) => {
        this.context.router.push("organizations/"+this.props.params.organizationID+"/applications/"+this.props.params.applicationID);
      });
    } else {
      NodeStore.createNodeKeys(this.props.params.devEUI, deviceKeys, (responseData) => {
        this.context.router.push("organizations/"+this.props.params.organizationID+"/applications/"+this.props.params.applicationID);
      });
    }
  }

  render() {
    return(
      <div>
        <div className="panel panel-default">
          <div className="panel-body">
            <DeviceKeysForm deviceKeys={this.state.deviceKeys} onSubmit={this.onSubmit} />
          </div>
        </div>
      </div>
    );
  }
}

export default NodeKeys;