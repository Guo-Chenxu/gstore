package com.guochenxu.gstore.entity;

import lombok.Data;

import java.util.List;
import java.util.Map;

@Data
public class GstoreResult {

    /**
     * 头部信息
     */
    @Data
    public class Head {
        private List<String> link;
        /**
         * 返回的查询结果的变量名
         */
        private List<String> vars;
    }

    @Data
    public class Results {
        private List<Map<String, Map<String, String>>> bindings;
    }

    private Head head;
    private Results results;
    private int StatusCode;
    private String StatusMsg;
    private int AnsNum;
    private int OutputLimit;
    private String ThreadId;
    private String QueryTime;
}