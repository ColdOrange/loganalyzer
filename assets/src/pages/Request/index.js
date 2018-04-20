// @flow

import * as React from 'react';
import { Row, Col } from 'antd';

import RequestMethod from './components/RequestMethod';
import HTTPVersion from './components/HTTPVersion';
import RequestURL from './components/RequestURL';
import StaticFile from './components/StaticFile';
import { fetchError } from 'utils/Modal';
import styles from './index.css';

class Request extends React.Component<{}> {
  error = false;

  // Handler error in parent component, in case it will show several times in child components
  errorHandler = () => {
    if (!this.error) {
      this.error = true;
      fetchError();
    }
  };

  render() {
    return (
      <div>
        <Row gutter={24}>
          <Col span={12}>
            <RequestMethod errorHandler={this.errorHandler} />
          </Col>
          <Col span={12}>
            <HTTPVersion errorHandler={this.errorHandler} />
          </Col>
        </Row>
        <div className={styles.divider} />
        <RequestURL errorHandler={this.errorHandler} />
        <div className={styles.divider} />
        <StaticFile errorHandler={this.errorHandler} />
      </div>
    );
  }
}

export default Request;
