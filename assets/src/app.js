// @flow

import * as React from 'react';
import { Layout, Menu, Icon } from 'antd';
import Summary from './components/summary'

import './app.css';

class App extends React.Component<{}> {
  render() {
    const { Header, Content, Footer, Sider } = Layout;
    return (
      <Layout>
        <Sider style={{ overflow: 'auto', height: '100vh', position: 'fixed', left: 0 }}>
          <div className="logo" />
          <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
            <Menu.Item key="1">
              <Icon type="home" />
              <span className="nav-text">Summary</span>
            </Menu.Item>
            <Menu.Item key="2">
              <Icon type="file-text" />
              <span className="nav-text">Page Views</span>
            </Menu.Item>
            <Menu.Item key="3">
              <Icon type="user" />
              <span className="nav-text">User Views</span>
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout style={{ marginLeft: 200 }}>
          <Header style={{ background: '#fff', padding: 0 }} />
          <Content style={{ margin: '24px 16px 0', overflow: 'initial', height: '100vh' }}>
            <div style={{ padding: 24, background: '#fff', textAlign: 'center', height: '100vh' }}>
              <Summary />
            </div>
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Log Analyzer ©2018 Created by Orange
          </Footer>
        </Layout>
      </Layout>
    );
  }
}

export default App;
