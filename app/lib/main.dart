import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:news_dashboard_app/pages/home/view.dart';
import 'package:get/get.dart';

String initialRoute = '/home';

void main() async {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ScreenUtilInit(
      designSize: const Size(1024, 768),
      minTextAdapt: true,
      splitScreenMode: true,
      builder: (context, child) {
        return GetMaterialApp(
          title: "News Dashboard",
          theme: ThemeData(
              primarySwatch: Colors.amber
          ),
          initialRoute: initialRoute,
          getPages: [
            GetPage(name: '/home', page: () => HomePage()),
          ],
        );
      },
    );
  }
}
