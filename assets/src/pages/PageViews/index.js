// @flow

import * as React from 'react';

import PageViewsDaily from './components/PageViewsDaily';
import PageViewsHourly from './components/PageViewsHourly';
import PageViewsMonthly from './components/PageViewsMonthly';
import styles from './index.css';

class PageViews extends React.Component<{}> {
  render() {
    return (
      <div>
        <PageViewsDaily />
        <div className={styles.divider} />
        <PageViewsHourly />
        <div className={styles.divider} />
        <PageViewsMonthly />
      </div>
    );
  }
}

export default PageViews;
