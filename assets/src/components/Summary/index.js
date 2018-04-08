// @flow

import * as React from 'react';
import { Card } from 'antd';

import KVCard from './components/KVCard';
import styles from './index.css';

type State = {
  fileName: string,
  timeRange: any, // TODO: how to deal with flow type of `Object` and `null`?
  pageViews: any,
  userViews: any,
  bandwidth: any,
  isLoaded: boolean,
}

class Summary extends React.Component<{}, State> {
  constructor() {
    super();
    this.state = {
      fileName: '',
      timeRange: null,
      pageViews: null,
      userViews: null,
      bandwidth: null,
      isLoaded: false,
    };
  }

  processData(data: Object) {
    // Time Range
    const timeRange = {
      'Start Time': data.start_time,
      'End Time': data.end_time,
    };
    // Page Views
    const durationMs = new Date(data.end_time) - new Date(data.start_time);
    const durationDay = durationMs / 1000 / 3600 / 24;
    const pageViews = {
      'Total Page Views': data.page_views,
      'Average Page Views per Day': Math.round(data.page_views / durationDay),
      'Average Page Views per User': Math.round(data.page_views / data.user_views),
    };
    // User Views
    const userViews = {
      'Total User Views': data.user_views,
      'Average User Views per Day': Math.round(data.user_views / durationDay),
    };
    // Bandwidth
    const bandwidth = {
      'Total Bandwidth': bandwidthToString(data.bandwidth),
      'Average Bandwidth per Day': bandwidthToString(data.bandwidth / durationDay),
      'Average Bandwidth per User': bandwidthToString(data.bandwidth / data.user_views),
    };
    this.setState({
      fileName: data.file_name,
      timeRange: timeRange,
      pageViews: pageViews,
      userViews: userViews,
      bandwidth: bandwidth,
      isLoaded: true,
    });
  }

  loadData() {
    fetch('/api/summary')
      .then(response => response.json())
      .then(  // TODO: handle error
        data => {
          this.processData(data);
        }
      );
  }

  componentDidMount() {
    this.loadData();
  }

  render() {
    const loading = !this.state.isLoaded;
    const title = <h4 className={styles.title}>{this.state.fileName}</h4>;
    return (
      <Card
        title={loading ? '' : title}
        className={styles.card}
        bordered={false}
      >
        <KVCard
          title="Time Range"
          loading={loading}
          data={mapToKVArray(this.state.timeRange)}
        />
        <KVCard
          title="Page Views"
          loading={loading}
          data={mapToKVArray(this.state.pageViews)}
        />
        <KVCard
          title="User Views"
          loading={loading}
          data={mapToKVArray(this.state.userViews)}
        />
        <KVCard
          title="Bandwidth"
          loading={loading}
          data={mapToKVArray(this.state.bandwidth)}
        />
      </Card>
    );
  }
}

// Convert bandwidth (in Byte) to human readable string
const bandwidthToString = (b: number): string => {
  if (b < 1024) {
    return b + ' B';
  }
  else if (b < 1024 * 1024) {
    return (b / 1024).toFixed(2) + ' KB';
  }
  else if (b < 1024 * 1024 * 1024) {
    return (b / 1024 / 1024).toFixed(2) + ' MB';
  }
  else {
    return (b / 1024 / 1024 / 1024).toFixed(2) + ' GB';
  }
};

// Map an object to a {key,value} object array
const mapToKVArray = (o: Object): Object[] => {
  return o == null ? [] : Object.entries(o).map(
    ([k, v]) => ({
      key: k,
      value: v,
    })
  );
};

export default Summary;
