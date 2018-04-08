// @flow

import * as React from 'react';
import { Card } from 'antd';

import KVTable from '../KVTable';

type Props = {
  title: string,
  loading: boolean,
  data: {
    key: string,
    value: any,
  }[],
}

class KVCard extends React.Component<Props> {
  render() {
    return (
      <Card
        title={this.props.title}
        loading={this.props.loading}
        type="inner"
        bodyStyle={{padding: 0}}
        bordered={false}
      >
        {
          this.props.loading ? '' : <KVTable data={this.props.data}/>
        }
      </Card>
    );
  }
}

export default KVCard;
