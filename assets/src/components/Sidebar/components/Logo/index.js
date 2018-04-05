// @flow

import * as React from 'react';

import styles from './index.css';

class Logo extends React.Component<{}> {
  render() {
    return (
      <div className={styles.logo}>
        <img
          className={styles.img}
          alt="logo"
          src="/static/images/logo.svg"
        />
        <span className={styles.span}>
          Log Analyzer
        </span>
      </div>
    );
  }
}

export default Logo;
