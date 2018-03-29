// @flow

import * as React from 'react';
import { Table } from 'antd';
import reqwest from 'reqwest';

import './index.css'

const columns = [{
    title: 'Start Time',
    dataIndex: 'start_time',
    width: '20%',
}, {
    title: 'End Time',
    dataIndex: 'end_time',
    width: '20%',
}, {
    title: 'Page Views',
    dataIndex: 'page_views',
    width: '20%',
}, {
    title: 'User Views',
    dataIndex: 'user_views',
    width: '20%',
}, {
    title: 'Bandwidth',
    dataIndex: 'bandwidth',
    width: '20%',
}];

class Summary extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            data: [],
            loading: false,
        };
    }

    fetch() {
        this.setState({ loading: true });
        reqwest({
            url: 'http://127.0.0.1:8080/api/summary',
            method: 'get',
            type: 'json'
        }).then((data) => {
            this.setState({
                loading: false,
                data: [data],
            });
        });
    }

    componentDidMount() {
        this.fetch();
    }

    render() {
        return (
            <Table columns={columns}
                   rowKey={record => record.registered}
                   dataSource={this.state.data}
                   loading={this.state.loading}
                   pagination={false}
            />
        );
    }
}

export default Summary;
