package com.example.emotiontracker.controller;

import jakarta.servlet.http.HttpSession;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@Controller
public class StatsController {

    @Value("${backend.base-url}")
    private String backendBaseUrl;

    @GetMapping("/stats")
    public String showStats(HttpSession session, Model model) {
        Object userId = session.getAttribute("user_id");
        Object userName = session.getAttribute("user_name");

        if (userId == null) {
            return "redirect:/login";
        } 

        
   // Integer userId = 1;
   // String userName = "测试用户";


        String url = backendBaseUrl + "/emotional/count/" + userId;
        RestTemplate restTemplate = new RestTemplate();

        try {
            ResponseEntity<Map> response = restTemplate.getForEntity(url, Map.class);
            if (response.getStatusCode() == HttpStatus.OK && response.getBody() != null) {
                Map data = (Map) response.getBody().get("data");
                model.addAttribute("good", data.get("good_num"));
                model.addAttribute("neutral", data.get("neutral_num"));
                model.addAttribute("bad", data.get("bad_num"));
            }
        } catch (Exception e) {
            model.addAttribute("error", "获取统计数据失败");
        }

        model.addAttribute("username", userName);
        return "stats";
    }
}
