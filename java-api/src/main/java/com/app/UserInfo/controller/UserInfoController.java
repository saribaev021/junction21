package com.app.UserInfo.controller;

import com.app.UserInfo.entity.UsersEntity;
import com.app.UserInfo.repository.UsersRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/user")
public class UserInfoController {
    @Autowired
    private UsersRepository usersRepository;

    @GetMapping("/{name}")
    public ResponseEntity getUserByName(@PathVariable String name) {
        try {
            return ResponseEntity.ok(usersRepository.findByName(name));
        } catch (Exception ex) {
            return ResponseEntity.badRequest().body("Error while searching user");
        }
    }

    @PostMapping("/add")
    public void saveUser(@RequestBody UsersEntity usersEntity) {
        try {
            usersRepository.save(usersEntity);
        } catch (Exception ex) {
            System.out.println("Error while saving user: \n" + ex.getMessage());
        }
    }
}
