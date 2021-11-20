package com.app.UserInfo.repository;

import com.app.UserInfo.entity.UsersEntity;
import org.springframework.data.repository.CrudRepository;

public interface UsersRepository extends CrudRepository<UsersEntity, Integer>{
    UsersEntity findByName(String name);
}
