// @flow

import * as React from 'react';
import { BrowserRouter, Route, Link, Switch } from 'react-router-dom';
import { Layout as AntLayout, Menu, Icon } from 'antd';

import Sidebar from './components/Sidebar';
import Page from './components/Page';

type Props = {
  pages: {
    link: string,
    icon: string,
    text: string,
    component: React.ComponentType<any>,
  }[],
  index: React.ComponentType<any>,
  notFound: React.ComponentType<any>,
}

type State = {
  collapsed: boolean,
}

class Layout extends React.Component<Props, State> {
  state = {
    collapsed: false,
  };

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  };

  render() {
    const MenuItems = this.props.pages.map(function (page) {
      return (
        <Menu.Item key={page.link}>
          <Link to={page.link}>
            <Icon type={page.icon} />
            <span className="nav-text">{page.text}</span>
          </Link>
        </Menu.Item>
      );
    });

    const Routes = this.props.pages.map(function (page) {
      return (
        <Route
          path={page.link}
          component={page.component}
          key={page.link}
        />
      );
    });
    Routes.unshift( // Index Page
      <Route
        exact path="/"
        component={this.props.index}
        key="/" />
    );
    Routes.push(    // Not Found
      <Route
        component={this.props.notFound}
        key="not-found"
      />
    );

    return (
      <BrowserRouter>
        <AntLayout>
          <Sidebar collapsed={this.state.collapsed}>
            {
              MenuItems
            }
          </Sidebar>
          <Page collapsed={this.state.collapsed} toggle={this.toggle}>
            <Switch>
              {
                Routes
              }
            </Switch>
          </Page>
        </AntLayout>
      </BrowserRouter>
    );
  }
}

export default Layout;
