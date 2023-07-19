package com.guochenxu.gstore;

import com.guochenxu.gstore.service.GstoreService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class GstoreApplicationTests {

    @Autowired
    GstoreService gs;

    @Test
    void contextLoads() {
    }

}
