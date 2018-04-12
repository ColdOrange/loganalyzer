// @flow

import * as React from 'react';

import ContentCard from 'components/ContentCard';
import KVCard from './components/KVCard';
import { fetchError } from 'utils/Modal';
import { bandwidthFormatter } from 'utils/BandwidthFormatter';

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
    const durationMs = new Date(data.end_time.replace(' ', 'T')) - new Date(data.start_time.replace(' ', 'T')); // replace... is workaround for safari
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
      'Total Bandwidth': bandwidthFormatter(data.bandwidth),
      'Average Bandwidth per Day': bandwidthFormatter(data.bandwidth / durationDay),
      'Average Bandwidth per User': bandwidthFormatter(data.bandwidth / data.user_views),
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
      .then(
        data => {
          if (data.status === 'failed') { // Server API error
            fetchError();
            console.log('Server API error');
          }
          else {
            this.processData(data);
          }
        },
        error => { // Fetch error
          fetchError();
          console.log(error);
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
