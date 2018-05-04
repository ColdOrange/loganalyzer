// @flow

import * as React from 'react';
import { BrowserRouter, Route, Redirect, Link, Switch } from 'react-router-dom';
import { Layout, Menu, Icon } from 'antd';

import Sidebar from 'components/Sidebar';
import Page from 'components/Page';
import Summary from 'pages/Summary';
import PageViews from 'pages/PageViews';
import UserViews from 'pages/UserViews';
import Bandwidth from 'pages/Bandwidth';
import Request from 'pages/Request';
import Response from 'pages/Response';
import UserAgent from 'pages/UserAgent';
import Referer from 'pages/Referer';
import NotFound from 'pages/NotFound';

type State = {
  collapsed: boolean,
}

class ReportLayout extends React.Component<{}, State> {
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
            <Menu.Item key="summary">
              <Link to="/report/summary">
                <Icon type="home" />
                <span className="nav-text">Summary</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="page-views">
              <Link to="/report/page-views">
                <Icon type="file-text" />
                <span className="nav-text">Page Views</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="user-views">
              <Link to="/report/user-views">
                <Icon type="user" />
                <span className="nav-text">User Views</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="bandwidth">
              <Link to="/report/bandwidth">
                <Icon type="cloud-download-o" />
                <span className="nav-text">Bandwidth</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="request">
              <Link to="/report/request">
                <Icon type="link" />
                <span className="nav-text">Request</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="response">
              <Link to="/report/response">
                <Icon type="export" />
                <span className="nav-text">Response</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="user-agent">
              <Link to="/report/user-agent">
                <Icon type="compass" />
                <span className="nav-text">User Agent</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="referer">
              <Link to="/report/referer">
                <Icon type="swap" />
                <span className="nav-text">Referer</span>
              </Link>
            </Menu.Item>
          </Sidebar>
          <Page collapsed={this.state.collapsed} toggle={this.toggle}>
            <Switch>
              <Route path="/report/summary" component={Summary} />
              <Route path="/report/page-views" component={PageViews} />
              <Route path="/report/user-views" component={UserViews} />
              <Route path="/report/bandwidth" component={Bandwidth} />
              <Route path="/report/request" component={Request} />
              <Route path="/report/response" component={Response} />
              <Route path="/report/user-agent" component={UserAgent} />
              <Route path="/report/referer" component={Referer} />
              <Route component={NotFound} />
              <Redirect from="/report" to="/report/summary" />
            </Switch>
          </Page>
        </Layout>
      </BrowserRouter>
    );
  }
}

export default ReportLayout;
