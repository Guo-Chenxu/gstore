package com.guochenxu.gstore.config;

import com.guochenxu.gstore.jgsc.GstoreConnector;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

/**
 * @program: gstore
 * @description: gstore 配置项
 * @author: 郭晨旭
 * @create: 2023-07-17 17:00
 * @version: 1.0
 **/

@Configuration
public class GstoreConfiguration {

    private static final String IP = "10.112.41.37";
    private static final Integer PORT = 9020;
    private static final String HTTP_TYPE = "ghttp";
    private static final String USER = "root";
    private static final String PASSWORD = "123456";


    @Bean
    public GstoreConnector gstoreConfig() {
        GstoreConnector gstoreConnector = new GstoreConnector(IP, PORT, HTTP_TYPE, USER, PASSWORD);
        return gstoreConnector;
    }
}
