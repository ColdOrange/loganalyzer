// @flow

import * as React from 'react';

import ContentCard from 'components/ContentCard';
import CustomPieChart from 'components/CustomPieChart';

type Props = {
  errorHandler: () => void,
}

type State = {
  data: {
    os: string,
    count: number,
  }[],
  isLoaded: boolean,
}

class OperatingSystem extends React.Component<Props, State> {
  state = {
    data: [],
    isLoaded: false,
  };

  loadData = () => {
    fetch('/api/user-agent/os')
      .then(response => response.json())
      .then(
        data => {
          if (data.status === 'failed') { // Server API error
            this.props.errorHandler();
            console.log('Server API error');
          }
          else {
            this.setState({
              data: data,
              isLoaded: true,
            });
          }
        },
        error => { // Fetch error
          this.props.errorHandler();
          console.log(error);
        }
      );
  };

  componentDidMount() {
    this.loadData();
  }

  render() {
    const loading = !this.state.isLoaded;
    const placeholder = // placeholder for rendering while loading
      <CustomPieChart
        data={[]}
        nameKey="os"
        dataKey="count"
      />;

    return (
      <ContentCard
        title="Operating System"
        loading={loading}
        placeholder={placeholder}
      >
        <CustomPieChart
          data={this.state.data}
          nameKey="os"
          dataKey="count"
        />
      </ContentCard>
    );
  }
}

export default OperatingSystem;
