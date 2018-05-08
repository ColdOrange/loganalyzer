// @flow

import * as React from 'react';
import { Link } from 'react-router-dom';
import { Table } from 'antd';

import ContentCard from 'components/ContentCard';
import { confirm, success, error as modalError, fetchError } from 'utils/Modal';

type State = {
  data: {
    id: number,
    file: string,
  }[],
  isLoaded: boolean,
}

class Reports extends React.Component<{}, State> { // TODO: selection
  state = {
    data: [],
    isLoaded: false,
  };

  columns = [
    {
      title: 'Id',
      dataIndex: 'id',
      width: '15%',
    },
    {
      title: 'Log File',
      dataIndex: 'file',
      width: '55%',
    },
    {
      title: 'Action',
      width: '15%',
      render: (text: any, record: Object) => (<Link to={`/reports/${record.id}`} target="_blank">View</Link>)
    },
    {
      title: '',
      width: '15%',
      render: (text: any, record: Object) => (<a href="#" onClick={() => this.deleteReport(record.id)}>Delete</a>)
    }
  ];

  deleteReport = (id: number) => {
    confirm({
      title: 'Are you sure to delete this report ?',
      content: 'The generated report and related data will be removed permanently. You can\'t undo this action.',
      onOk: () => {
        fetch(`/api/reports/${id}`, {
          method: 'DELETE',
        })
          .then(response => response.json())
          .then(
            data => {
              if (data.status === 'successful') {
                success({
                  title: 'Success',
                  content: 'Delete report successfully!',
                });
              }
              else {
                modalError({
                  title: 'Error',
                  content: 'Delete report failed, please again try later.',
                });
              }
            },
            error => { // Fetch error
              modalError({
                title: 'Error',
                content: 'Delete report failed, please again try later.',
              });
              console.log(error);
            }
          );
      },
    });
  };

  loadData = () => {
    fetch('/api/reports')
      .then(response => response.json())
      .then(
        data => {
          if (data && data.status === 'failed') { // Server API error // TODO: data == null
            fetchError();
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
      <ContentCard title="Reports">
        <Table
          dataSource={this.state.data}
          columns={this.columns}
          loading={loading}
          rowKey="id"
        />
      </ContentCard>
    );
  }
}

export default Reports;
