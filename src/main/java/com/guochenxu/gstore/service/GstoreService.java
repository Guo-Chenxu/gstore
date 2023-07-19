package com.guochenxu.gstore.service;

import com.guochenxu.gstore.entity.GstoreResult;

/**
 * @program: gstore
 * @description: gstore服务接口
 * @author: 郭晨旭
 * @create: 2023-07-17 17:22
 * @version: 1.0
 **/
public interface GstoreService {

    GstoreResult query(String database, String sparQL);
}
