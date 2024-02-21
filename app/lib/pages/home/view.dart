import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import 'logic.dart';

class HomePage extends StatelessWidget {
  HomePage({Key? key}) : super(key: key);

  final logic = Get.put(HomeLogic());
  final state = Get.find<HomeLogic>().state;

  Widget buildClock(){
    return Obx(() => Column(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          state.currentTime.value,
          style: TextStyle(
            color: Colors.white,
            fontSize: 80.sp,
            fontWeight: FontWeight.bold,
          ),
        ),
      ],
    ));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        width: MediaQuery.of(context).size.width,
        height: MediaQuery.of(context).size.height,
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomCenter,
            colors: [
              Color(0xff3f7cf5),
              Color(0xff91caff),
            ],
          ),
        ),
        child: Stack(
          children: [
            Row(
              children: [
                Container(
                  margin: EdgeInsets.only(left: 10.w, top: 10.w,),
                  width: 390.w,
                  height: 369.h,
                  decoration: BoxDecoration(
                    color: Colors.white.withOpacity(0.3),
                    borderRadius: BorderRadius.circular(10),
                  ),
                  child: buildClock(),
                ),
                Container(
                  margin: EdgeInsets.only(left: 10.w, top: 10.w,),
                  width: 300.w,
                  height: 369.h,
                  decoration: BoxDecoration(
                    color: Colors.white.withOpacity(0.3),
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
