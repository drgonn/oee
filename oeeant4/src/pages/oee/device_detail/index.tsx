import React, {useEffect, useRef, useState} from 'react';
import {PageHeaderWrapper} from '@ant-design/pro-layout';
import {
  Chart,
  Interval,
  Axis,
  Tooltip,
  Coordinate,
  Legend,
  View,
  Annotation,
} from "bizcharts";
import { DataView } from '@antv/data-set';
import DataSet from "@antv/data-set";
import { queryWorktimeList, updateWorktime , addWorktime , removeWorktime  } from "../worktime/service"

import ProCard from '@ant-design/pro-card';
import {Card, Descriptions} from 'antd';
import ProTable from '@ant-design/pro-table';
import moment from 'moment';
import { GridContent } from '@ant-design/pro-layout';


const operationTabList = [

  {
    key: 'WorkTime',
    tab: '工作时间分析',
  },
  {
    key: 'DeviceDetail',
    tab: '驾动率',
  },

  {
    key: 'Failure',
    tab: '故障率',
  },
];

const data = [
  // { name: 'London', 月份: 'Jan.', 月均降雨量: 18.9 },
  // { name: 'London', 月份: 'Feb.', 月均降雨量: 28.8 },
  {name: "良品", amount: 3, time: "03:02:26"},
  {name: "良品", amount: 3, time: "19:02:24"},
  {name: "良品", amount: 3, time: "19:02:25"},
 {name: "良品", amount: 33, time: "2020-09-16 19:02:26"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:27"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:28"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:29"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:30"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:31"},
 {name: "良品", amount: 3, time: "2020-09-16 19:02:32"},
{name: "良品", amount: 3, time: "2020-09-16 19:02:33"},
];


const modify_worktime = (data) => {
  let chart_data = []
  for (let o of data) {
    let type1
    let name1
    switch (o.type) {
      case 1:
        type1 = "停止时间";
        name1 = "休息时间";
        break
      case 2:
        type1 = "负荷时间";
        name1 = "日常管理时间";
        break
      case 3:
        name1 = "停机时间";
        type1 = "负荷时间";
        break
      case 4:
        type1 = "负荷时间";
        name1 = "运转时间";
        break
      case 5:
        name1 = "计划停止时间";
        type1 = "负荷时间";
        break
      case 6:
        type1 = "负荷时间";
        name1 = "日常管理时间";
        break
    }
    chart_data.push({ type: type1, value: o.seconds, name:name1,good:o.good,amount:o.amount});
  }
  return chart_data;
}

const cadv = (data) => {
  // 通过 DataSet 计算百分比
  const chartdata = modify_worktime(data)
  const dv = new DataView();
  console.log(chartdata)
  dv.source(chartdata).transform({
    type: 'percent',
    field: 'value',
    dimension: 'type',
    as: 'percent',
  });
  return dv
}
const cadv1 = (data) => {
  // 通过 DataSet 计算百分比
  const chartdata = modify_worktime(data)
  const dv1 = new DataView();
  console.log(chartdata)
  dv1.source(chartdata).transform({
    type: 'percent',
    field: 'value',
    dimension: 'name',
    as: 'percent',
  });
  return dv1
}
const chartdata = (data) => {
  let chart_data = [];
  for (let o of data) {
    chart_data.push({ name: "总量", amount: o.amount, time: o.end_time.substring(10,20) });
    chart_data.push({ name: "良品", amount: o.good, time: o.end_time.substring(10,20) });
    // chart_data.push({ month: o.trade_date, city: "评分", temperature: o.score });
  }
  return chart_data;
};

export default (e) => {
  const [worktime, setworktime] = useState([]);
  const [statedv, setdv    ] = useState([]);
  const [statedv1, setdv1    ] = useState([]);
  const [statedv2, setdv2    ] = useState([]);
  const [tabkey, setTabkey] = useState('WorkTime');
  const { id } = e.location.query;
  const onOperationTabChange = (key) => {
    setTabkey(key);
  };
  useEffect(() => {
    console.log(e.location)
    async function fetchData() {
      const wtime = await queryWorktimeList({device_id:id,pageSize:700});
      console.log(wtime)
      setworktime(wtime.data);
      setdv(cadv(wtime.data).rows)
      setdv1(cadv1(wtime.data).rows)
      setdv2(chartdata(wtime.data))


      // setDetail(data);
      // const list = [];
      // const flowlist = await queryFlowdata({ iccid: data.iccid, month });
      // if (flowlist.length) {
      //   flowlist.reverse().map((item) => {
      //     list.push({
      //       x: moment(item.date).format('MM-DD'),
      //       y: item.dataUsage,
      //     });
      //   });
      // }
      // setFlowdata(list);
    }
    fetchData();
  }, []);

  // console.log(statedv.slice(1))
  // console.log(statedv.slice(1).pop())
  // console.log(statedv.length>0 && statedv.slice(1).pop().value)

  useEffect(() => {
    console.log(statedv2)
  }, []);


  const contentList = {
    DeviceDetail: (
      <Descriptions bordered>
        <Descriptions.Item label="编号"> {worktime.length} </Descriptions.Item>
        <Descriptions.Item label="总流量"> ddd M </Descriptions.Item>
        <Descriptions.Item label="剩余量"> dd M </Descriptions.Item>
        <Descriptions.Item label="本月总量">sdd M </Descriptions.Item>
      </Descriptions>
    ),
    WorkTime: (
    <ProCard style={{ marginTop: 3 }}  ghost>
      <ProCard bordered layout="center" colSpan={12}>
        <Descriptions bordered>
          <Descriptions.Item label="停止时间" span={3}> {statedv.length>1 && statedv[1].value}s </Descriptions.Item>
          <Descriptions.Item label="休息时间" span={3}> {statedv.length>1 && statedv[1].value}s </Descriptions.Item>
          <Descriptions.Item label="负荷时间" span={3}> {statedv.length>1 && statedv[0].value}s </Descriptions.Item>
          <Descriptions.Item label="运转时间" span={1}>{statedv1.length>1 && statedv1[2].value}s</Descriptions.Item>
          <Descriptions.Item label="日常管理" span={1}>{statedv1.length>1 && statedv1[1].value}s</Descriptions.Item>
          <Descriptions.Item label="停机时间" span={1}>{statedv1.length>1 && statedv1[0].value}s</Descriptions.Item>
        </Descriptions>
      </ProCard>

      <ProCard colSpan="40%" bordered>
        <Chart
          height={400}
          data={statedv}
          autoFit
          scale={{
            percent: {
              formatter: (val) => {
                val = `${(val * 100).toFixed(2)}%`;
                return val;
              },
            }
          }}
        >
          <Coordinate type="theta" radius={0.5} />
          <Axis visible={false} />
          <Legend visible={false} />
          <Tooltip showTitle={false} />
          <Interval
            position="percent"
            adjust="stack"
            color="type"
            element-highlight
            style={{
              lineWidth: 1,
              stroke: '#fff',
            }}
            label={['type', {
              offset: -15,
            }]}
          />
          <View data={statedv1}>
            <Coordinate type="theta" radius={0.75} innerRadius={0.5 / 0.75} />
            <Interval
              position="percent"
              adjust="stack"
              color={['name', ['#BAE7FF', '#7FC9FE', '#71E3E3', '#ABF5F5', '#8EE0A1', '#BAF5C4']]}
              element-highlight
              style={{
                lineWidth: 1,
                stroke: '#fff',
              }}
              label="name"
            />
          </View>
        </Chart>
      </ProCard>
    </ProCard>
    ),
  Failure: (
    <>
  <ProCard style={{ marginTop: 3 }}  ghost>
    <ProCard bordered layout="center" colSpan="60%">
      <Descriptions bordered>
        <Descriptions.Item label="良品数量" span={3}> {statedv.length>1 && statedv[1].value}个 </Descriptions.Item>
        <Descriptions.Item label="故障率" span={3}> {statedv.length>1 && statedv[1].value} % </Descriptions.Item>

      </Descriptions>
    </ProCard>

    <ProCard colSpan="40%" bordered>
      <Chart
        height={400}
        data={statedv}
        autoFit
        scale={{
          percent: {
            formatter: (val) => {
              val = `${(val * 100).toFixed(2)}%`;
              return val;
            },
          }
        }}
      >
        <Coordinate type="theta" radius={0.5} />
        <Axis visible={false} />
        <Legend visible={false} />
        <Tooltip showTitle={false} />
        <Interval
          position="percent"
          adjust="stack"
          color="type"
          element-highlight
          style={{
            lineWidth: 1,
            stroke: '#fff',
          }}
          label={['type', {
            offset: -15,
          }]}
        />
      </Chart>
    </ProCard>

  </ProCard>
  <ProCard colSpan="100%" bordered>
    <Chart height={500} padding="auto"
           data={statedv2.length > 1 && statedv2}
           // data={data}
           autoFit>
      <Interval
        adjust={[
          {
            type: 'stack',
          },
        ]}
        color="name"
        position="time*amount"
      />
      <Tooltip shared />
    </Chart>
  </ProCard>
  </>
    ),
  };

  return (
    <GridContent>

    </GridContent>

    // <PageHeaderWrapper>
    //   <Card
    //     tabList={operationTabList}
    //     onTabChange={(key) => {
    //       onOperationTabChange(key);
    //     }}
    //   >
    //     {contentList[tabkey]}
    //   </Card>
    // </PageHeaderWrapper>
  );
};
