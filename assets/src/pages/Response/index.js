// @flow

import * as React from 'react';
import { Row, Col } from 'antd';

import StatusCode from './components/StatusCode';
import ResponseTime from './components/ResponseTime';
import ResponseURL from './components/ResponseURL';
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
            <StatusCode errorHandler={this.errorHandler} />
          </Col>
          <Col span={12}>
            <ResponseTime errorHandler={this.errorHandler} />
          </Col>
        </Row>
        <div className={styles.divider} />
        <ResponseURL errorHandler={this.errorHandler} />
      </div>
    );
  }
}

export default Request;
