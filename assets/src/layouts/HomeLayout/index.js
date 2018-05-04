// @flow

import * as React from 'react';
import { BrowserRouter, Route, Link, Switch } from 'react-router-dom';
import { Layout, Menu, Icon } from 'antd';

import Sidebar from 'components/Sidebar';
import Page from 'components/Page';
import Home from 'pages/Home';
import Config from 'pages/Config';
import Reports from 'pages/Reports';
import NotFound from 'pages/NotFound';

type State = {
  collapsed: boolean,
}

class ConfigLayout extends React.Component<{}, State> {
  state = {
    collapsed: false,
  };

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  };

  render() {
    return (
      <BrowserRouter>
        <Layout>
          <Sidebar collapsed={this.state.collapsed}>
            <Menu.Item key="home">
              <Link to="/">
                <Icon type="home" />
                <span className="nav-text">Home</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="config">
              <Link to="/config">
                <Icon type="setting" />
                <span className="nav-text">Config</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="reports">
              <Link to="/reports">
                <Icon type="copy" />
                <span className="nav-text">Reports</span>
              </Link>
            </Menu.Item>
          </Sidebar>
          <Page collapsed={this.state.collapsed} toggle={this.toggle}>
            <Switch>
              <Route path="/" exact component={Home} />
              <Route path="/config" component={Config} />
              <Route path="/reports" component={Reports} />
              <Route component={NotFound} />
            </Switch>
          </Page>
        </Layout>
      </BrowserRouter>
    );
  }
}

export default ConfigLayout;
