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

    /**
     * 具体的查询结果
     */
    @Data
    public class SOP {
        String type;
        String dataType;
        String value;
    }

    @Data
    public class Bindings {
        /**
         * 具体的查询结果, map 中的键是 head.var 里面的值, map 中的值是具体的查询结果
         */
        private Map<String, SOP> map;
    }

    @Data
    public class Results {
        private List<String> bindings;
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