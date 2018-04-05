// @flow

import * as React from 'react';
import { Layout } from 'antd';

import styles from './index.css';

const { Header, Content, Footer } = Layout;

type Props = {
  children?: React.Node,
}

class Page extends React.Component<Props> {
  render() {
    return (
      <Layout className={styles.page}>
        <Header className={styles.header} />
        <Content className={styles.content}>
          {
            this.props.children
          }
        </Content>
        <Footer className={styles.footer}>Log Analyzer © 2018 Created by Orange</Footer>
      </Layout>
    );
  }
}

export default Page;
