package com.example.emotiontracker.controller;

import jakarta.servlet.http.HttpSession;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@Controller
public class CheckinController {  

    @Value("${backend.base-url}")
    private String backendBaseUrl;

    @GetMapping("/checkin")
    public String showCheckinPage(HttpSession session, Model model) {
        Object username = session.getAttribute("user_name");
        Object userIdObj = session.getAttribute("user_id");

        if (username == null || userIdObj == null) {
            return "redirect:/login";
        }

        model.addAttribute("username", username);
        return "checkin";
    }

    @PostMapping("/checkin")
    public String doCheckin(HttpSession session, Model model) {
        Object userIdObj = session.getAttribute("user_id");
        Object username = session.getAttribute("user_name");

        if (userIdObj == null) {
            return "redirect:/login";
        }

        Long userId = Long.parseLong(String.valueOf(userIdObj));
        String url = backendBaseUrl + "/checkin/" + userId;

        RestTemplate restTemplate = new RestTemplate();

        try {
            ResponseEntity<Map> response = restTemplate.postForEntity(url, null, Map.class);
            Map<String, Object> body = response.getBody();

            if (body != null) {
                if (body.containsKey("error")) {
                    model.addAttribute("result", body.get("error"));
                } else if (body.containsKey("msg")) {
                    model.addAttribute("result", body.get("msg"));
                } else if (body.containsKey("data")) {
                    model.addAttribute("result", body.get("data"));
                } else {
                    model.addAttribute("result", "未收到打卡结果");
                }
            }

        } catch (Exception e) {
            model.addAttribute("result", "打卡异常，请稍后再试");
        }

        model.addAttribute("username", username);
        return "checkin";
    }
}
