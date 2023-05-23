    YunUsers yunUsers = userFacade.getYunUsersByUserId(userId) ;
        Properties properties = new Properties() ;
        InputStream resourceAsStream = getClass().getClassLoader().getResourceAsStream("dchealth.properties");
        properties.load(resourceAsStream);
        String newPassword =properties.getProperty("newPassword");
        if("".equals(newPassword)||null==newPassword){
            newPassword = "123456" ;
