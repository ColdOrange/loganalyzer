// @flow

import * as React from 'react';
import { BrowserRouter, Route, Link, Switch } from 'react-router-dom';
import { Layout, Menu, Icon } from 'antd';
import 'antd/dist/antd.css';

import Sidebar from './components/Sidebar';
import Page from './components/Page';
import Summary from './pages/Summary';
import PageViews from './pages/PageViews';
import UserViews from './pages/UserViews';
import Bandwidth from './pages/Bandwidth';
import Request from './pages/Request';
import Response from './pages/Response';
import UserAgent from './pages/UserAgent';
import Referer from './pages/Referer';
import NotFound from './pages/NotFound';

class App extends React.Component<{}> {
  render() {
    return (
      <BrowserRouter>
        <Layout>
          <Sidebar>
            <Menu.Item key="summary">
              <Link to="summary">
                <Icon type="home" />
                <span className="nav-text">Summary</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="page-views">
              <Link to="page-views">
                <Icon type="file-text" />
                <span className="nav-text">Page Views</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="user-views">
              <Link to="user-views">
                <Icon type="user" />
                <span className="nav-text">User Views</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="bandwidth">
              <Link to="bandwidth">
                <Icon type="cloud-download-o" />
                <span className="nav-text">Bandwidth</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="request">
              <Link to="request">
                <Icon type="link" />
                <span className="nav-text">Request</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="response">
              <Link to="response">
                <Icon type="export" />
                <span className="nav-text">Response</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="user-agent">
              <Link to="user-agent">
                <Icon type="compass" />
                <span className="nav-text">User Agent</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="referer">
              <Link to="referer">
                <Icon type="swap" />
                <span className="nav-text">Referer</span>
              </Link>
            </Menu.Item>
          </Sidebar>
          <Page>
            <Switch>
              <Route exact path="/" component={Summary} />
              <Route path="/summary" component={Summary} />
              <Route path="/page-views" component={PageViews} />
              <Route path="/user-views" component={UserViews} />
              <Route path="/bandwidth" component={Bandwidth} />
              <Route path="/request" component={Request} />
              <Route path="/response" component={Response} />
              <Route path="/user-agent" component={UserAgent} />
              <Route path="/referer" component={Referer} />
              <Route component={NotFound} />
            </Switch>
          </Page>
        </Layout>
      </BrowserRouter>
    );
  }
}

export default App;
