// @flow

import * as React from 'react';
import { BrowserRouter, Route, Link, Switch } from 'react-router-dom';
import { Layout, Menu, Icon } from 'antd';
import 'antd/dist/antd.css';

import Sidebar from './components/Sidebar';
import Page from './components/Page';
import Summary from  './components/Summary';
import PageViews from  './components/PageViews';
import UserViews from  './components/UserViews';
import NotFound from  './components/NotFound';

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
          </Sidebar>
          <Page>
            <Switch>
              <Route exact path="/" component={Summary} />
              <Route path="/summary" component={Summary} />
              <Route path="/page-views" component={PageViews} />
              <Route path="/user-views" component={UserViews} />
              <Route component={NotFound} />
            </Switch>
          </Page>
        </Layout>
      </BrowserRouter>
    );
  }
}

export default App;
