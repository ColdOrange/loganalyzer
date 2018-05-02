// @flow

import * as React from 'react';
import { Row, Col } from 'antd';

import OperatingSystem from './components/OperatingSystem';
import Device from './components/Device';
import Browser from './components/Browser';
import { fetchError } from 'utils/Modal';
import styles from './index.css';

class UserAgent extends React.Component<{}> {
  error = false;

  // Handle error in parent component, in case it will show several times in children
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
            <OperatingSystem errorHandler={this.errorHandler} />
          </Col>
          <Col span={12}>
            <Device errorHandler={this.errorHandler} />
          </Col>
        </Row>
        <div className={styles.divider} />
        <Browser errorHandler={this.errorHandler} />
      </div>
    );
  }
}

export default UserAgent;
