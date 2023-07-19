package com.guochenxu.gstore.service.impl;

import com.alibaba.fastjson2.JSON;
import com.guochenxu.gstore.entity.GstoreResult;
import com.guochenxu.gstore.jgsc.GstoreConnector;
import com.guochenxu.gstore.service.GstoreService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @program: gstore
 * @description: gstore服务实现类
 * @author: 郭晨旭
 * @create: 2023-07-17 17:23
 * @version: 1.0
 **/

@Service
public class GstoreServiceImpl implements GstoreService {

    @Autowired
    private GstoreConnector gc;

    @Override
    public GstoreResult query(String database, String sparQL) {
        String result = gc.query(database, "json", sparQL);
        GstoreResult gstore = JSON.parseObject(result, GstoreResult.class);
        return gstore;
    }
}
