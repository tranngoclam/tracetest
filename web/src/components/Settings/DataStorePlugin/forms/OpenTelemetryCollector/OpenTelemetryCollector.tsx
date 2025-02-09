import {Col, Row} from 'antd';
import Ingestor from './Ingestor';
import Configuration from './Configuration';

const OpenTelemetryCollector = () => {
  return (
    <>
      <Row gutter={[16, 16]}>
        <Col span={16}>
          <Ingestor />
        </Col>
      </Row>
      <Row gutter={[16, 16]}>
        <Col span={16}>
          <Configuration />
        </Col>
      </Row>
    </>
  );
};

export default OpenTelemetryCollector;
