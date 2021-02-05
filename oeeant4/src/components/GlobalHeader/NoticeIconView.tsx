import { connect, ConnectProps } from 'umi';
import { Tag, message } from 'antd';
import groupBy from 'lodash/groupBy';
import moment from 'moment';
import { NoticeItem } from '@/models/global';
import { CurrentUser } from '@/models/user';
import { ConnectState } from '@/models/connect';
import React, { useState, useRef, useEffect} from 'react';
import NoticeIcon from '../NoticeIcon';
import styles from './index.less';
import { queryMessageList, updateMessage , addMessage , removeMessage  } from './service';
import {ActionType} from "@ant-design/pro-table";
import {PageHeaderWrapper} from "@ant-design/pro-layout";

export interface GlobalHeaderRightProps extends Partial<ConnectProps> {
  notices?: NoticeItem[];
  currentUser?: CurrentUser;
  fetchingNotices?: boolean;
  onNoticeVisibleChange?: (visible: boolean) => void;
  onNoticeClear?: (tabName?: string) => void;
}



const GlobalHeaderRight = () => {
  const [messagelist, setMessagelist] = useState([]);
  const actionRef = useRef<ActionType>();



  const getMessage = async(obj)=> {
    const data = await queryMessageList(obj)
    if (data.success){
      console.log("接收消息",data)
      setMessagelist(data.data)
    }
  };
  const readMessage = async(obj)=> {
    const data = await updateMessage(obj)
    if (data.success){
      console.log("接息",data)
      getMessage({read:0})
    }
  };
  const changeReadState = (clickedItem: NoticeItem): void => {
    const { id } = clickedItem;
    console.log(clickedItem)
    readMessage({read:1,id:id})
  };

  const handleNoticeClear = (title: string, key: string) => {


    message.success(`${'清空了'} ${title}`);

  };

  const getNoticeData = (): {
    [key: string]: NoticeItem[];
  } => {
    const { notices = [] } = this.props;

    if (!notices || notices.length === 0 || !Array.isArray(notices)) {
      return {};
    }

    const newNotices = notices.map((notice) => {
      const newNotice = { ...notice };

      if (newNotice.datetime) {
        newNotice.datetime = moment(notice.datetime as string).fromNow();
      }

      if (newNotice.id) {
        newNotice.key = newNotice.id;
      }

      if (newNotice.extra && newNotice.status) {
        const color = {
          todo: '',
          processing: 'blue',
          urgent: 'red',
          doing: 'gold',
        }[newNotice.status];
        newNotice.extra = (
          <Tag
            color={color}
            style={{
              marginRight: 0,
            }}
          >
            {newNotice.extra}
          </Tag>
        );
      }

      return newNotice;
    });
    return groupBy(newNotices, 'type');
  };

  const getUnreadData = (noticeData: { [key: string]: NoticeItem[] }) => {
    const unreadMsg: {
      [key: string]: number;
    } = {};
    Object.keys(noticeData).forEach((key) => {
      const value = noticeData[key];

      if (!unreadMsg[key]) {
        unreadMsg[key] = 0;
      }

      if (Array.isArray(value)) {
        unreadMsg[key] = value.filter((item) => !item.read).length;
      }
    });
    return unreadMsg;
  };

  // render() {
  //   const { currentUser, fetchingNotices, onNoticeVisibleChange } = this.props;
  //   // const noticeData = this.getNoticeData();

  //   const {message} = this.state
  //   console.log(message)
  //   // const noticeData = message
  //   console.log(noticeData)
  //   console.log(unreadMsg)
  //   console.log(this.state)

  const noticeData = {message:[
      {id:"123",title:"ddddd",description:"ffffffffff",status:"todo"},
    ]};
  const unreadMsg = getUnreadData(noticeData);

  useEffect(()=>{
    console.log("初始化")
    getMessage({read:0})
    console.log(messagelist)
  },[])

  return (
    <NoticeIcon
      actionRef={actionRef}
      className={styles.action}
        count={messagelist.length}
        onItemClick={(item) => {
          changeReadState(item as NoticeItem);
        }}
        loading={getMessage}
        clearText="清空"
        viewMoreText="查看更多"
        onClear={handleNoticeClear}
        onPopupVisibleChange={true}
        onViewMore={() => message.info('Click on view more')}
        clearClose
      >
        <NoticeIcon.Tab
          tabKey="message"
          count={messagelist.length}
          list={messagelist}
          title="消息"
          emptyText="您已读完所有消息"
          showViewMore
        />
      </NoticeIcon>
    );
}

export default GlobalHeaderRight;
