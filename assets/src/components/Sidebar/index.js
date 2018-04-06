// @flow

import * as React from 'react';
import { Layout, Menu } from 'antd';

import Logo from './components/Logo';
import styles from './index.css';

const { Sider } = Layout;

type Props = {
  children?: React.Node,
}

class Sidebar extends React.Component<Props> {
  render() {
    const currentPage = window.location.pathname.substring(1);
    return (
      <Sider className={styles.sider}>
        <Logo />
        <Menu className={styles.menu} theme="light" mode="inline" defaultSelectedKeys={currentPage}>
          {
            this.props.children
          }
        </Menu>
      </Sider>
    );
  }
}

export default Sidebar;
