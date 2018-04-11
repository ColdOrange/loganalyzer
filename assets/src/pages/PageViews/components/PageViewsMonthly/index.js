// @flow

import * as React from 'react';

import ContentCard from 'components/ContentCard';
import CustomLineChart from 'components/CustomLineChart';
import styles from './index.css';

type State = {
  data: {
    time: string,
    pv: number,
  }[],
  isLoaded: boolean,
}

class PageViewsMonthly extends React.Component<{}, State> {
  state = {
    data: [],
    isLoaded: false,
  };

  loadData = () => {
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
  };

  componentDidMount() {
    this.loadData();
  }

  render() {
    const loading = !this.state.isLoaded;
    const placeholder = // placeholder for rendering while loading
      <CustomLineChart
        data={[]}
        xAxisKey="time"
        lineKey="pv"
        color="#8884d8"
      />;

    return (
      <ContentCard
        title="Monthly"
        loading={loading}
        placeholder={placeholder}
      >
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
