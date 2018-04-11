// @flow

import * as React from 'react';

import ContentCard from 'components/ContentCard';
import KVCard from './components/KVCard';

type State = {
  fileName: string,
  timeRange: { [string]: any },
  pageViews: { [string]: any },
  userViews: { [string]: any },
  bandwidth: { [string]: any },
  isLoaded: boolean,
}

class Summary extends React.Component<{}, State> {
  state = {
    fileName: '',
    timeRange: {},
    pageViews: {},
    userViews: {},
    bandwidth: {},
    isLoaded: false,
  };

  processData = (data: Object) => {
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
  };

  loadData = () => {
    fetch('/api/summary')
      .then(response => response.json())
      .then(  // TODO: handle error
        data => { // TODO: handle server api error (status: failed)
          this.processData(data);
        }
      );
  };

  componentDidMount() {
    this.loadData();
  }

  render() {
    const loading = !this.state.isLoaded;
    return (
      <ContentCard title={this.state.fileName}>
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
      </ContentCard>
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
