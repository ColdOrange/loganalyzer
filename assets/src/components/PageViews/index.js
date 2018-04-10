// @flow

import * as React from 'react';

import PageViewsDaily from './components/PageViewsDaily';
import PageViewsHourly from './components/PageViewsHourly';
import PageViewsMonthly from './components/PageViewsMonthly';

class PageViews extends React.Component<{}> {
  render() {
    return (
      <div>
        <PageViewsDaily />
        <div style={{paddingTop: 24}} />
        <PageViewsHourly />
        <div style={{paddingTop: 24}} />
        <PageViewsMonthly />
      </div>
    );
  }
}

export default PageViews;
