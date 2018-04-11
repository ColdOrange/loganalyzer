// @flow

import * as React from 'react';
import { Layout, Menu } from 'antd';

import Logo from './components/Logo';
import styles from './index.css';

type Props = {
  children?: React.Node,
}

class Sidebar extends React.Component<Props> {
  render() {
    const { Sider } = Layout;
    let currentPage = window.location.pathname.substring(1);
    if (currentPage === '') {
      currentPage = 'summary';
    }
    return (
      <Sider className={styles.sider}>
        <Logo />
        <Menu
          defaultSelectedKeys={[currentPage]}
          className={styles.menu}
          theme="light"
          mode="inline"
        >
          {
            this.props.children
          }
        </Menu>
      </Sider>
    );
  }
}

export default Sidebar;
