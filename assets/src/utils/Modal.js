// @flow

import * as React from 'react';
import { Modal } from 'antd';

type ButtonClickHandler = (event: SyntheticEvent<HTMLButtonElement>) => void;

export const error = (title: string | React.Node, content: string | React.Node, onOK?: ButtonClickHandler) => {
  Modal.error({
    title: title,
    content: content,
    onOK: onOK,
  });
};

export const fetchError = () => {
  error('Fetch data error', 'Sorry. There seems to be an error when fetching data from server. Please try again later.');
};
