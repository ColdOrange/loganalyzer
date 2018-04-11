// @flow

import * as React from 'react';
import { Card, Spin } from 'antd';

import styles from './index.css';

type Props = {
  title: string,
  loading?: boolean,
  placeholder?: React.Node, // TODO: better way to render loading indicator?
  children?: React.Node,
}

class ContentCard extends React.Component<Props> {
  render() {
    const title = <h4 className={styles.title}>{this.props.title}</h4>;
    const { loading = false } = this.props;
    return (
      <Card
        title={title}
        className={styles.card}
        bordered={false}
      >
        <Spin spinning={loading}>
          {
            loading ? this.props.placeholder : this.props.children
          }
        </Spin>
      </Card>
    );
  }
}

export default ContentCard;
