// @flow

import * as React from 'react';
import 'antd/dist/antd.css';

import Layout from './layout';
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
      <Layout
        pages={[
          {
            link: '/summary',
            icon: 'home',
            text: 'Summary',
            component: Summary,
          },
          {
            link: '/page-views',
            icon: 'file-text',
            text: 'Page Views',
            component: PageViews,
          },
          {
            link: '/user-views',
            icon: 'user',
            text: 'User Views',
            component: UserViews,
          },
          {
            link: '/bandwidth',
            icon: 'cloud-download-o',
            text: 'Bandwidth',
            component: Bandwidth,
          },
          {
            link: '/request',
            icon: 'link',
            text: 'Request',
            component: Request,
          },
          {
            link: '/response',
            icon: 'export',
            text: 'Response',
            component: Response,
          },
          {
            link: '/user-agent',
            icon: 'compass',
            text: 'User Agent',
            component: UserAgent,
          },
          {
            link: '/referer',
            icon: 'swap',
            text: 'Referer',
            component: Referer,
          },
        ]}
        index={Summary}
        notFound={NotFound}
      />
    );
  }
}

export default App;
