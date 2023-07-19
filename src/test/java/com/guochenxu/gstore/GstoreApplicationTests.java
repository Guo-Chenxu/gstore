package com.guochenxu.gstore;

import com.guochenxu.gstore.service.GstoreService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.concurrent.*;

@SpringBootTest
class GstoreApplicationTests {

    @Autowired
    GstoreService gs;

    @Test
    void contextLoads() {
    }

    int count = 0;

    @Test
    void testTimer() {
        ExecutorService executor = Executors.newSingleThreadExecutor();
        Future<String> future = executor.submit(new Callable<String>() {
            public String call() throws Exception {
                // 执行耗时操作
                Thread.sleep(1000);
                return "Hello, World!";
            }
        });

        try {
            // 等待操作完成，最多等待3秒
            String result = future.get(3000, TimeUnit.MILLISECONDS);
            System.out.println(result);
        } catch (InterruptedException | ExecutionException e) {
            e.printStackTrace();
        } catch (TimeoutException e) {
            // 如果等待超时，则取消任务的执行
            future.cancel(true);
            System.out.println("Task timed out!");
        }

        executor.shutdown();
    }
}
