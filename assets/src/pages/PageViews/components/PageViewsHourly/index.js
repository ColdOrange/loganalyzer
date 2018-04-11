// @flow

import * as React from 'react';
import { Select } from 'antd';

import ContentCard from 'components/ContentCard';
import CustomLineChart from 'components/CustomLineChart';
import styles from './index.css';

type State = {
  days: string[],
  data: {
    time: string,
    pv: number,
  }[],
  isLoaded: boolean,
}

class PageViewsHourly extends React.Component<{}, State> {
  state = {
    days: [],
    data: [],
    isLoaded: false,
  };

  loadData = (date: string) => {
    fetch(`/api/page-views/hourly?date=${date}`)
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
    fetch('/api/page-views/daily')
      .then(response => response.json())
      .then(
        data => {
          this.setState({
            days: data.map(item => item.time)
          });
          this.loadData(this.state.days[0]);
        }
      );
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
        title="Hourly"
        loading={loading}
        placeholder={placeholder}
      >
        <Select
          defaultValue={this.state.days[0]}
          className={styles.select}
          onChange={this.loadData}
        >
          {
            this.state.days.map(date =>
              <Select.Option key={date} value={date}>
                {date}
              </Select.Option>
            )
          }
        </Select>
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

export default PageViewsHourly;
