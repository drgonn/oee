import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import {Col, Row, Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, List ,Card , Upload, Typography, Tabs} from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';
import { getToken } from '@/utils/authority';
import { routerRedux } from 'dva/router';
import { EditOutlined, EllipsisOutlined, SettingOutlined } from '@ant-design/icons';
import {
	G2,
	Chart,
	Interval,
	Interaction
} from "bizcharts";
import { TableListItem } from './data.d';
import { queryDeviceList, updateDevice , addDevice , removeDevice  } from './service';
import {CardListItemDataType} from "../../../../../../../front/src/pages/list/card-list/data";
import { FormattedMessage,history, Link } from 'umi';
const { Meta } = Card;
const { TextArea } = Input;
const { Paragraph } = Typography;
const { TabPane } = Tabs;
import { DatePicker, Space } from 'antd';
import { PieChart } from 'bizcharts';


import { SmileTwoTone, BulbOutlined, CheckCircleTwoTone } from '@ant-design/icons';

const { RangePicker } = DatePicker;


const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeDevice({
      ids: selectedRows.map((row) => row.id),
    });
    if (res.success){
    hide();
    message.success('删除成功，即将刷新');
    return true;
        }else{
          message.error(res.errmsg||'请求失败请重试！');
          hide();
          return;
        }
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const TableList: React.FC<{}> = () => {
  const [createModalVisible, handleModalVisible] = useState<boolean>(false);
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [filename,setFilename]=useState('')
  const [formvalues,setValues]=useState({})
  const [devicelist,setdevice]=useState([])
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '编号',
      dataIndex: 'sn',
      valueType: 'text',
    },
    {
      title: '名称',
      dataIndex: 'name',
      render: (val, record) => <Link to={`/oee/device/detail?id=${record.id}`}>{val}</Link>,
    },
    {
      title: '类型',
      dataIndex: 'type',
      valueEnum: {
        1: { text:'三轴'},
        2: { text:'五轴'},
      }
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => (
        <>
          <a
            onClick={() => {
              handleUpdateModalVisible(true);
              setValues(record);
              setId(record.id);
            }}
          >
            修改
          </a>
          <Divider type="vertical" />
        </>
      ),
    },
  ];

  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateDevice({...values})
        if(res.success){
          hide()
          message.success('修改成功！')
          handleUpdateModalVisible(false);
          actionRef.current.reload();
        }else{
          message.error(res.errmsg||'请求失败请重试！');
          hide();
          return;
        }
      }catch(error){
        // message.error('error');
        hide();
      }
    })
  }
  const handleAdd = ()=>{
    form
      .validateFields().then(async(values)=>{
      const hide=message.loading('正在提交...')
      try{
        const res=await addDevice({...values, img: filename})
        if(res.success){
          hide()
          message.success('创建成功！')
          handleModalVisible(false);
          queryDevice({pageindex:1,pageSize:1000})
          actionRef.current.reload();
        }else{
          message.error(res.errmsg||'请求失败请重试1！');
          hide();
          return;
        }
      }catch(error){
        // message.error('请求失败请重试2！');
        hide();
      }
    })
  }
  const queryDevice=async(obj)=>{
    const data = await queryDeviceList(obj);
    if(data.success){
      if(data.data&& data.data.length){
        data.data.map((item,index)=>{
          item.key=index
        })
      }
      setdevice(data.data)
    }
  };
  useEffect(() => {
    queryDevice({pageindex:1,pageSize:1000})
  }, []);


  const props = {
    name: 'file',
    action: `/api/upload?token=${getToken()}`,
    headers: {
      authorization: 'authorization-text',
    },
    showUploadList:false,
    listType:"text",
    onChange(info) {
      if (info.file.status !== 'uploading') {
        console.log(info.file, info.fileList);
      }
      console.log(info.file.response)
      if(info.file.response && info.file.response.success===false){
        message.error(info.file.response.errmsg ||'上传失败')
        return
      } else {
        info.file.response && setFilename(info.file.response.upload_file)
      }
      if (info.file.status === 'done') {
        message.success(`${info.file.name} 上传成功`);
      } else if (info.file.status === 'error') {
        message.error(`${info.file.name} 文件上传失败.`);
      }
    },
  };

  const data = [
	{ name: 'London', month: 'Jan.', avgRainfall: 18.9 },
	{ name: 'London', month: 'Feb.', avgRainfall: 28.8 },
	{ name: 'London', month: 'Mar.', avgRainfall: 39.3 },
	{ name: 'London', month: 'Apr.', avgRainfall: 20.4 },
	{ name: 'London', month: 'May', avgRainfall: 47 },
	{ name: 'London', month: 'Jun.', avgRainfall: 20.3 },
	{ name: 'London', month: 'Jul.', avgRainfall: 24 },
	{ name: 'London', month: 'Aug.', avgRainfall: null },
	{ name: 'Paris', month: 'Jan.', avgRainfall: 19.9 },
	{ name: 'Paris', month: 'Feb.', avgRainfall: 28.8 },
	{ name: 'Paris', month: 'Mar.', avgRainfall: 29.3 },
	{ name: 'Paris', month: 'Apr.', avgRainfall: null },
	{ name: 'Paris', month: 'May', avgRainfall: 67 },
	{ name: 'Paris', month: 'Jun.', avgRainfall: null },
	{ name: 'Paris', month: 'Jul.', avgRainfall: 23 },
	{ name: 'Paris', month: 'Aug.', avgRainfall: null },
	{ name: 'Berlin', month: 'Jan.', avgRainfall: 12.4 },
	{ name: 'Berlin', month: 'Feb.', avgRainfall: 23.2 },
	{ name: 'Berlin', month: 'Mar.', avgRainfall: 34.5 },
	{ name: 'Berlin', month: 'Apr.', avgRainfall: null },
	{ name: 'Berlin', month: 'May', avgRainfall: 52.6 },
	{ name: 'Berlin', month: 'Jun.', avgRainfall: 35.5 },
	{ name: 'Berlin', month: 'Jul.', avgRainfall: 37.4 },
	{ name: 'Berlin', month: 'Aug.', avgRainfall: 42.4 },
];

  const data1 = [
    {
      type: '机台一',
      value: 27,
    },
    {
      type: '机台二',
      value: 25,
    },
    {
      type: '机台一',
      value: 18,
    },

    {
      type: '机台二',
      value: 10,
    },
    {
      type: '机台三',
      value: 5,
    },
  ];
  return (
    <PageHeaderWrapper
      extra={[
        <Button key="1" type="primary" onClick={() => handleModalVisible(true)}>
          <PlusOutlined /> 新建机台
        </Button>,
      ]}
      actionRef={actionRef}
    >
      <List<Partial<CardListItemDataType>>
        rowKey="id"
        grid={{
          gutter: 16,
          xs: 1,
          sm: 2,
          md: 3,
          lg: 3,
          xl: 4,
          xxl: 5,
        }}
        dataSource = {devicelist}
        renderItem={(item) => {
          if (item && item.id ) {
            console.log(item)
            return (
              <List.Item key={item.id} >
                <Card
                  title={item.name}
                  hoverable
                  bodyStyle={{ padding: '20px 24px 8px 24px' }}
                  style={{ high: 240 }}
                  cover={<img alt="example" src={item.img_url}
                              height="200" width="200"
                              onClick={() => {
                                history.push(`/oee/machine/detail?id=${item.id}`);
                              }}

                  />}
                  actions={[
                    <SettingOutlined key="setting"
                                     onClick={() => {
                                       setValues( item );
                                       handleUpdateModalVisible(true);
                                       setId(item.id);
                                     }
                                     }
                    />,
                    <BulbOutlined   style={{ color: 'green' }}/>

                  ]}
                >
                  <Card.Meta
                    // avatar={<img alt="example" src={item.img_url}
                    //             onClick={() => {
                    //               history.push(`/oee/machine/detail?id=${item.id}`);
                    //             }}
                    //
                    // />}
                    // avatar=
                    //   {<img alt="" className={styles.cardAvatar} src={item.avatar} />}
                    title={<a>{item.sn}</a>}
                    // description={
                    //   <Paragraph className={styles.item} ellipsis={{ rows: 3 }}>
                    //     {item.description}
                    //   </Paragraph>
                    // }
                  />

                </Card>
              </List.Item>
            );
          }
        }}
      />

        <Card  bordered={false} >
            <Tabs
              // tabBarExtraContent={<Button>Extra Action</Button>}
              tabBarExtraContent={
                <div >
                  <div >
                    <Button >
                      今日
                    </Button>
                    <Button >
                      本周
                    </Button>
                    <Button >
                      本月
                    </Button>
                    <RangePicker
                      // onChange={handleRangePickerChange}
                      style={{ width: 256 }}
                    />
                  </div>

                </div>
              }
              size="large"
              tabbarstyle={{ marginbottom: 24 }}
            >
              <TabPane tab="产量" key="1">
                <Row>
                  <Col xl={18} lg={12} md={12} sm={24} xs={24}>
                    <div >
                      <Chart height={400} padding="auto" data={data} autoFit filter={[
                        ['avgRainfall', val => val != null] // 图表将会只渲染过滤后的数据
                      ]}>
                        <Interval
                          adjust={[
                            {
                              type: 'dodge',
                              marginRatio: 0,
                            },
                          ]}
                          color="name"
                          position="month*avgRainfall"
                        />
                        <Interaction type="active-region" />
                      </Chart>

                    </div>
                  </Col>
                  <Col xl={6} lg={12} md={12} sm={24} xs={24}>
                    <div>


    <PieChart
      data={data1}
      title={{
        visible: true,
        text: '产量-各个机台产量汇总',
      }}
      description={{
        visible: true,
        text: '产量百分比',
      }}
      radius={0.8}
      angleField='value'
      colorField='type'
      label={{
        visible: true,
        type: 'outer',
        offset: 20,
      }}
    />
                    </div>
                  </Col>
                </Row>
              </TabPane>
        <TabPane tab="良品率" key="2">
          良品率
        </TabPane>
            </Tabs>
        </Card>
      <Modal
        title="新建机器"
        // visible={createModalVisible}
        onOk={handleAdd}
        // onCancel={() => handleModalVisible(false)}
      >
        <Form form={form} initialValues={formvalues}>
          <Form.Item
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器类型"
          >
            <Select>
            <Option value={1}>三轴</Option>
            <Option value={2}>五轴</Option>
            </Select>
          </Form.Item>

          <Form.Item
            name='ip'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器ip地址"
          >
            <Input placeholder="请输入ip地址" />
          </Form.Item>
          <Form.Item
            // name='file'
            label="机器图片"
          >
            <Upload {...props}>
              <Button>Click to Upload</Button>
            </Upload>
          </Form.Item>

        </Form>
      </Modal>
      <Modal
        title="编辑机器"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false); }}
      >
        <Form form={form} initialValues={formvalues}>
          <Form.Item
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器类型"
          >
            <Select>
            <Option value={1}>三轴</Option>
            <Option value={2}>五轴</Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='ip'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器ip地址"
          >
            <Input placeholder="请输入ip地址" />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
