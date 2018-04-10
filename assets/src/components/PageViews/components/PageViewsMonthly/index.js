// @flow

import * as React from 'react';

import ContentCard from '../../../ContentCard';
import CustomLineChart from '../../../CustomLineChart';
import styles from './index.css';

type State = {
  data: {
    time: string,
    pv: number,
  }[],
  isLoaded: boolean,
}

class PageViewsMonthly extends React.Component<{}, State> {
  constructor() {
    super();
    this.state = {
      data: [],
      isLoaded: false,
    };
  }

  loadData() {
    fetch('/api/page-views/monthly')
      .then(response => response.json())
      .then(  // TODO: handle error
        data => { // TODO: handle server api error (status: failed)
          this.setState({
            data: data,
            isLoaded: true,
          });
        }
      );
  }

  componentDidMount() {
    this.loadData();
  }

  render() {
    // const loading = !this.state.isLoaded; // TODO: unused
    return (
      <ContentCard title="Monthly">
        <div className={styles.container}>
          <CustomLineChart
            data={this.state.data}
            xAxisKey="time"
            lineKey="pv"
            color="#8884d8"
          />
        </div>
      </ContentCard>
    );
  }
}

export default PageViewsMonthly;
