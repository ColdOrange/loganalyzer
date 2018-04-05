// @flow

import * as React from 'react';
import { Table } from 'antd';

import styles from './index.css';

const columns = [
  {
    title: 'Start Time',
    dataIndex: 'start_time',
    width: '20%',
  },
  {
    title: 'End Time',
    dataIndex: 'end_time',
    width: '20%',
  },
  {
    title: 'Page Views',
    dataIndex: 'page_views',
    width: '20%',
  },
  {
    title: 'User Views',
    dataIndex: 'user_views',
    width: '20%',
  },
  {
    title: 'Bandwidth',
    dataIndex: 'bandwidth',
    width: '20%',
  }
];

type State = {
  data: any[],
  loading: boolean,
}

class Summary extends React.Component<{}, State> {
  constructor() {
    super();
    this.state = {
      data: [],
      loading: false,
    };
  }

  loadData() {
    this.setState({loading: true});
    fetch('/api/summary')
      .then(res => res.json())
      .then(  // TODO: handle error
        result => {
          this.setState({
            data: [result],
            loading: false,
          });
        }
      );
  }

  componentDidMount() {
    this.loadData();
  }

  render() {
    return (
      <div className={styles.summary}>
        <Table
          columns={columns}
          dataSource={this.state.data}
          loading={this.state.loading}
          rowKey="start_time"
          pagination={false}
        />
      </div>
    );
  }
}

export default Summary;
