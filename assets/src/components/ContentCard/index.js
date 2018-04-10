// @flow

import * as React from 'react';
import { Card } from 'antd';

import styles from './index.css';

type Props = {
  title: string,
  loading?: boolean,
  children?: React.Node,
}

class ContentCard extends React.Component<Props> {
  render() {
    const title = <h4 className={styles.title}>{this.props.title}</h4>;
    const { loading = false } = this.props;
    return (
      <Card
        title={title}
        loading={loading}
        className={styles.card}
        bordered={false}
      >
        {
          this.props.children
        }
      </Card>
    );
  }
}

export default ContentCard;
