// @flow

import * as React from 'react';
import { Table } from 'antd';

import ContentCard from 'components/ContentCard';

const columns = [
  {
    title: 'Request URL',
    dataIndex: 'requestURL',
    width: '60%',
  },
  {
    title: 'Count',
    dataIndex: 'count',
    width: '40%',
  }
];

type Props = {
  errorHandler: () => void,
}

type State = {
  data: {
    requestURL: string,
    count: number,
  }[],
  isLoaded: boolean,
}

class RequestURL extends React.Component<Props, State> {
  state = {
    data: [],
    isLoaded: false,
  };

  loadData = () => {
    fetch('/api/request-url')
      .then(response => response.json())
      .then(
        data => {
          if (data.status === 'failed') { // Server API error
            this.props.errorHandler();
            console.log('Server API error');
          }
          else {
            this.setState({
              data: data,
              isLoaded: true,
            });
          }
        },
        error => { // Fetch error
          this.props.errorHandler();
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
      <ContentCard title="Request URL">
        <Table
          dataSource={this.state.data}
          columns={columns}
          loading={loading}
          rowKey="requestURL"
        />
      </ContentCard>
    );
  }
}

export default RequestURL;
