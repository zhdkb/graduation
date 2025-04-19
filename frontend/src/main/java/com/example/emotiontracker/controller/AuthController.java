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
public class AuthController {

    @Value("${backend.base-url}")
    private String backendBaseUrl;

    // ========== 登录页面 ==========
    @GetMapping("/login")
    public String showLoginPage() {
        return "login";
    }

    @PostMapping("/login")
    public String doLogin(@RequestParam String username,
                          @RequestParam String password,
                          HttpSession session,
                          Model model) {

        String url = backendBaseUrl + "/login";
        RestTemplate restTemplate = new RestTemplate();

        Map<String, String> reqBody = new HashMap<>();
        reqBody.put("username", username);
        reqBody.put("password", password);

        try {
            ResponseEntity<Map> response = restTemplate.postForEntity(url, reqBody, Map.class);

            System.out.println("后端返回：" + response.getBody());

            if (response.getStatusCode() == HttpStatus.OK && response.getBody() != null) {
                Map<String, Object> body = response.getBody();
                Object dataObj = body.get("data");

                if (dataObj instanceof Map data) {
                    session.setAttribute("user_id", data.get("user_id"));
                    session.setAttribute("user_name", data.get("user_name"));
                    session.setAttribute("token", data.get("token"));
                    return "redirect:/home";
                }
            }

            model.addAttribute("error", "用户名或密码错误，请重试");
        } catch (Exception e) {
            System.out.println("登录异常：" + e.getMessage());
            model.addAttribute("error", "登录失败，服务器无响应");
        }

        return "login";
    }

    // ========== 注册页面 ==========
    @GetMapping("/signup")
    public String showRegisterPage() {
        return "signup";
    }

    @PostMapping("/signup")
    public String doRegister(@RequestParam String username,
                             @RequestParam String password,
                             @RequestParam String re_password,
                             Model model) {

        String url = backendBaseUrl + "/signup"; // ✅ 正确注册接口
        RestTemplate restTemplate = new RestTemplate();

        Map<String, String> reqBody = new HashMap<>();
        reqBody.put("username", username);
        reqBody.put("password", password);
        reqBody.put("re_password", re_password);

        try {
            ResponseEntity<Map> response = restTemplate.postForEntity(url, reqBody, Map.class);

            System.out.println("注册接口响应：" + response.getBody());

            if (response.getStatusCode() == HttpStatus.OK && response.getBody() != null) {
                Object msg = response.getBody().get("msg");
                if ("success".equals(msg)) {
                    return "redirect:/login";
                } else {
                    model.addAttribute("error", "注册失败：" + msg);
                }
            }
        } catch (Exception e) {
            model.addAttribute("error", "注册失败，请检查网络连接或用户名是否重复");
        }

        return "signup";
    }

    // ========== 登出 ==========
    @GetMapping("/logout")
    public String logout(HttpSession session) {
        session.invalidate();
        return "redirect:/login";
    }
}
