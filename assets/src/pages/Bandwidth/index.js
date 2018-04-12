// @flow

import * as React from 'react';

import BandwidthDaily from './components/BandwidthDaily';
import BandwidthHourly from './components/BandwidthHourly';
import BandwidthMonthly from './components/BandwidthMonthly';
import { fetchError } from 'utils/Modal';
import styles from './index.css';

class Bandwidth extends React.Component<{}> {
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
        <BandwidthDaily errorHandler={this.errorHandler} />
        <div className={styles.divider} />
        <BandwidthHourly errorHandler={this.errorHandler} />
        <div className={styles.divider} />
        <BandwidthMonthly errorHandler={this.errorHandler} />
      </div>
    );
  }
}

export default Bandwidth;
