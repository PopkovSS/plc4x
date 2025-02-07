/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by code-generation. DO NOT EDIT.

namespace org.apache.plc4net.drivers.mqtt.readwrite.model
{

    public enum MQTT_RetainHandling
    {
        SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_THE_SUBSCRIBE = 0x0,
        SEND_RETAINED_MESSAGES_AT_SUBSCRIBE_ONLY_IF_THE_SUBSCRIPTION_DOES_NOT_CURRENTLY_EXIST = 0x1,
        DO_NOT_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_SUBSCRIBE = 0x2,
    }

}

