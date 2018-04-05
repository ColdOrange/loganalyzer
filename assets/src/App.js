// @flow

import * as React from 'react';
import { Layout, Menu, Icon } from 'antd';
import 'antd/dist/antd.css';

import Sidebar from './components/Sidebar';
import Page from './components/Page';
import Summary from  './components/Summary';

class App extends React.Component<{}> {
  render() {
    return (
      <Layout>
        <Sidebar>
          <Menu.Item key="summary">
            <Icon type="home" />
            <span className="nav-text">Summary</span>
          </Menu.Item>
          <Menu.Item key="page-views">
            <Icon type="file-text" />
            <span className="nav-text">Page Views</span>
          </Menu.Item>
          <Menu.Item key="user-views">
            <Icon type="user" />
            <span className="nav-text">User Views</span>
          </Menu.Item>
        </Sidebar>
        <Page>
          <Summary />
        </Page>
      </Layout>
    );
  }
}

export default App;
