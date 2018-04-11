// @flow

import * as React from 'react';

import PageViewsDaily from './components/PageViewsDaily';
import PageViewsHourly from './components/PageViewsHourly';
import PageViewsMonthly from './components/PageViewsMonthly';
import { fetchError } from 'utils/Modal';
import styles from './index.css';

class PageViews extends React.Component<{}> {
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
        <PageViewsDaily errorHandler={this.errorHandler} />
        <div className={styles.divider} />
        <PageViewsHourly errorHandler={this.errorHandler} />
        <div className={styles.divider} />
        <PageViewsMonthly errorHandler={this.errorHandler} />
      </div>
    );
  }
}

export default PageViews;
