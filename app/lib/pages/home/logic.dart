import 'dart:async';

import 'package:get/get.dart';
import 'package:intl/intl.dart';

import 'state.dart';

class HomeLogic extends GetxController {
  final HomeState state = HomeState();
  late Timer _timer;

  @override
  void onInit() {
    super.onInit();
    _startTimer();
  }

  @override
  void onClose() {
    super.onClose();
    // 释放资源
    _timer.cancel();
  }

  void _startTimer() {
    // 每秒更新一次时间
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      state.currentTime.value = DateFormat('HH:mm:ss').format(DateTime.now());
    });
  }
}
