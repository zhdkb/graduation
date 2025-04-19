package com.example.emotiontracker.controller;

import jakarta.servlet.http.HttpSession;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.util.HashMap;
import java.util.Map;

@Controller
public class HomeController {

    @Value("${backend.base-url}")
    private String backendBaseUrl;

    // 展示首页
    @GetMapping("/home")
    public String showHomePage(HttpSession session, Model model) {
        Object username = session.getAttribute("user_name");

        if (username == null) {
            return "redirect:/login";
        }

        model.addAttribute("username", username);
        return "home";
    }

    // 情绪分析功能
    @PostMapping("/analyze")
    public String analyzeEmotion(@RequestParam String text,
                                 HttpSession session,
                                 Model model) {

        Object userIdObj = session.getAttribute("user_id");
        Object token = session.getAttribute("token");

        if (userIdObj == null || token == null) {
            return "redirect:/login";
        }

        
        Long userId;
        try {
            userId = Long.parseLong(String.valueOf(userIdObj));
        } catch (NumberFormatException e) {
            model.addAttribute("error", "用户 ID 异常，请重新登录");
            return "home";
        }

        // 准备请求
        String url = backendBaseUrl + "/emotional";
        RestTemplate restTemplate = new RestTemplate();

        Map<String, Object> request = new HashMap<>();
        request.put("user_id", userId);
        request.put("text", text);

        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        HttpEntity<Map<String, Object>> entity = new HttpEntity<>(request, headers);

        try {
            ResponseEntity<Map> response = restTemplate.postForEntity(url, entity, Map.class);

            System.out.println("后端响应：" + response.getBody());

            if (response.getStatusCode() == HttpStatus.OK && response.getBody() != null) {
                Object dataObj = response.getBody().get("data");

                if (dataObj instanceof Map<?, ?> data) {
                    model.addAttribute("sentiment", data.get("sentiment"));
                    model.addAttribute("sentimentType", data.get("sentiment_type"));
                    model.addAttribute("analyzedText", data.get("text"));
                } else {
                    model.addAttribute("error", "后端未能返回情绪分析结果");
                }
            } else {
                model.addAttribute("error", "后端服务未响应，请稍后再试");
            }
        } catch (Exception e) {
            System.out.println("情绪分析异常：" + e.getMessage());
            model.addAttribute("error", "情绪分析失败，请稍后重试");
        }

        model.addAttribute("username", session.getAttribute("user_name"));
        return "home";
    }
}
