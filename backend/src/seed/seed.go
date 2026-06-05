package seed

import (
	"time"

	"noise-complaint-backend/src/database"
	"noise-complaint-backend/src/models"

	"gorm.io/gorm"
)

func SeedAll() {
	db := database.GetDB()

	seedTimePeriods(db)
	seedComplaints(db)
	seedSamplingPoints(db)
	seedNoiseReadings(db)
	seedEvidenceAttachments(db)
	seedRectificationMeasures(db)
	seedRetestRecords(db)
	seedRuleConfigs(db)
	seedAnomalyEvents(db)
}

func seedTimePeriods(db *gorm.DB) {
	periods := []models.TimePeriod{
		{
			PeriodNo:    "TP-2024-001",
			PeriodName:  "昼间时段(6:00-22:00)",
			PeriodType:  models.TimePeriodTypeDay,
			TimeFrom:    "06:00",
			TimeTo:      "22:00",
			DayLimit:    65.0,
			NightLimit:  55.0,
			Status:      models.TimePeriodStatusActive,
			Description: "居民区昼间噪声标准",
		},
		{
			PeriodNo:    "TP-2024-002",
			PeriodName:  "夜间时段(22:00-6:00)",
			PeriodType:  models.TimePeriodTypeNight,
			TimeFrom:    "22:00",
			TimeTo:      "06:00",
			DayLimit:    55.0,
			NightLimit:  45.0,
			Status:      models.TimePeriodStatusActive,
			Description: "居民区夜间噪声标准",
		},
		{
			PeriodNo:    "TP-2024-003",
			PeriodName:  "工业区昼间时段",
			PeriodType:  models.TimePeriodTypeDay,
			TimeFrom:    "06:00",
			TimeTo:      "22:00",
			DayLimit:    70.0,
			NightLimit:  60.0,
			Status:      models.TimePeriodStatusActive,
			Description: "工业区昼间噪声标准",
		},
	}
	db.Create(&periods)
}

func seedComplaints(db *gorm.DB) {
	complaints := []models.Complaint{
		{
			ComplaintNo:     "CMP-2024-0001",
			Title:           "华泰小区夜间施工噪声扰民",
			Description:     "华泰小区3号楼居民投诉，夜间22:00后仍有施工噪声",
			Status:          models.ComplaintStatusPending,
			Level:           models.ComplaintLevelHigh,
			ComplainantName: "张明",
			ComplainantTel:  "138****5678",
			EnterpriseName:  "华泰建筑有限公司",
			EnterpriseAddr:  "朝阳区望京街道华泰小区",
			ResponsibleUser: "李处理员",
			BatchNo:         "BATCH-2024-001",
		},
		{
			ComplaintNo:     "CMP-2024-0002",
			Title:           "工业园区风机噪声超标",
			Description:     "工业园区风机运行噪声超过标准限值",
			Status:          models.ComplaintStatusReviewing,
			Level:           models.ComplaintLevelMedium,
			ComplainantName: "王芳",
			ComplainantTel:  "139****1234",
			EnterpriseName:  "东方机械制造厂",
			EnterpriseAddr:  "海淀区上地工业园区A栋",
			ResponsibleUser: "赵处理员",
			BatchNo:         "BATCH-2024-001",
		},
		{
			ComplaintNo:     "CMP-2024-0003",
			Title:           "商业广场空调外机噪声",
			Description:     "商业广场空调外机噪声影响周边居民休息",
			Status:          models.ComplaintStatusConfirmed,
			Level:           models.ComplaintLevelLow,
			ComplainantName: "刘洋",
			ComplainantTel:  "137****9876",
			EnterpriseName:  "万达商业广场",
			EnterpriseAddr:  "西城区金融街万达广场",
			ResponsibleUser: "孙处理员",
			BatchNo:         "BATCH-2024-002",
		},
		{
			ComplaintNo:     "CMP-2024-0004",
			Title:           "学校周边交通噪声严重",
			Description:     "学校周边交通噪声严重影响教学秩序",
			Status:          models.ComplaintStatusSupplement,
			Level:           models.ComplaintLevelUrgent,
			ComplainantName: "陈校长",
			ComplainantTel:  "136****4567",
			EnterpriseName:  "市教育局",
			EnterpriseAddr:  "东城区和平里街道第一中学",
			ResponsibleUser: "周处理员",
			BatchNo:         "BATCH-2024-002",
		},
		{
			ComplaintNo:     "CMP-2024-0005",
			Title:           "餐饮店排风扇噪声扰民",
			Description:     "餐饮店排风扇运行噪声影响楼上居民",
			Status:          models.ComplaintStatusArchived,
			Level:           models.ComplaintLevelMedium,
			ComplainantName: "吴女士",
			ComplainantTel:  "135****7890",
			EnterpriseName:  "老北京炸酱面馆",
			EnterpriseAddr:  "丰台区方庄小区芳群园",
			ResponsibleUser: "郑处理员",
			BatchNo:         "BATCH-2024-003",
		},
		{
			ComplaintNo:     "CMP-2024-0006",
			Title:           "地铁站通风口噪声投诉",
			Description:     "地铁站通风口设备噪声超标",
			Status:          models.ComplaintStatusRejected,
			Level:           models.ComplaintLevelLow,
			ComplainantName: "马先生",
			ComplainantTel:  "134****2345",
			EnterpriseName:  "地铁运营公司",
			EnterpriseAddr:  "朝阳区国贸地铁站",
			ResponsibleUser: "钱处理员",
			BatchNo:         "BATCH-2024-003",
			Remark:          "经核实噪声未超标，投诉不成立",
		},
		{
			ComplaintNo:     "CMP-2024-0007",
			Title:           "医院周边救护车警报噪声",
			Description:     "医院救护车警报声频繁影响周边居民",
			Status:          models.ComplaintStatusPending,
			Level:           models.ComplaintLevelMedium,
			ComplainantName: "黄女士",
			ComplainantTel:  "133****6789",
			EnterpriseName:  "市第一人民医院",
			EnterpriseAddr:  "宣武区白纸坊街道",
			ResponsibleUser: "李处理员",
			BatchNo:         "BATCH-2024-004",
		},
		{
			ComplaintNo:     "CMP-2024-0008",
			Title:           "建筑工地打桩机噪声",
			Description:     "建筑工地打桩机作业噪声严重超标",
			Status:          models.ComplaintStatusReviewing,
			Level:           models.ComplaintLevelUrgent,
			ComplainantName: "林先生",
			ComplainantTel:  "132****0123",
			EnterpriseName:  "中建三局",
			EnterpriseAddr:  "通州区梨园镇建筑工地",
			ResponsibleUser: "赵处理员",
			BatchNo:         "BATCH-2024-004",
		},
	}
	db.Create(&complaints)
}

func seedSamplingPoints(db *gorm.DB) {
	points := []models.SamplingPoint{
		{
			PointNo:         "SP-2024-0001",
			PointName:       "华泰小区3号楼东侧",
			Address:         "朝阳区望京街道华泰小区3号楼东侧围墙外1米",
			Longitude:       116.4756,
			Latitude:        39.9872,
			Status:          models.SamplingPointStatusCompleted,
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			PointNo:         "SP-2024-0002",
			PointName:       "工业园区A栋东侧",
			Address:         "海淀区上地工业园区A栋东侧围墙外1米",
			Longitude:       116.3124,
			Latitude:        40.0532,
			Status:          models.SamplingPointStatusSampling,
			ComplaintID:     2,
			ComplaintNo:     "CMP-2024-0002",
			ResponsibleUser: "采样员刘明",
			BatchNo:         "BATCH-2024-001",
		},
		{
			PointNo:         "SP-2024-0003",
			PointName:       "万达广场北侧居民楼",
			Address:         "西城区金融街万达广场北侧居民楼3层阳台",
			Longitude:       116.3621,
			Latitude:        39.9123,
			Status:          models.SamplingPointStatusCompleted,
			ComplaintID:     3,
			ComplaintNo:     "CMP-2024-0003",
			ResponsibleUser: "采样员张华",
			BatchNo:         "BATCH-2024-002",
		},
		{
			PointNo:         "SP-2024-0004",
			PointName:       "第一中学教学楼北侧",
			Address:         "东城区和平里街道第一中学教学楼北侧围墙外1米",
			Longitude:       116.4215,
			Latitude:        39.9567,
			Status:          models.SamplingPointStatusScheduled,
			ComplaintID:     4,
			ComplaintNo:     "CMP-2024-0004",
			ResponsibleUser: "采样员李军",
			BatchNo:         "BATCH-2024-002",
		},
		{
			PointNo:         "SP-2024-0005",
			PointName:       "芳群园小区2号楼",
			Address:         "丰台区方庄小区芳群园2号楼3层阳台",
			Longitude:       116.4356,
			Latitude:        39.8721,
			Status:          models.SamplingPointStatusArchived,
			ComplaintID:     5,
			ComplaintNo:     "CMP-2024-0005",
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-003",
		},
	}
	db.Create(&points)
}

func seedNoiseReadings(db *gorm.DB) {
	readings := []models.NoiseReading{
		{
			ReadingNo:       "NR-2024-0001",
			SamplingPointID: 1,
			PointNo:         "SP-2024-0001",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			TimePeriodID:    2,
			PeriodName:      "夜间时段(22:00-6:00)",
			MeasurementDate: time.Date(2024, 3, 15, 0, 0, 0, 0, time.Local),
			MeasurementTime: "23:30",
			Leq:             68.5,
			Lmax:            75.2,
			Lmin:            52.3,
			L10:             71.2,
			L90:             55.6,
			StandardLimit:   55.0,
			ExceedValue:     13.5,
			IsExceeded:      true,
			Status:          models.NoiseReadingStatusConfirmed,
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			ReadingNo:       "NR-2024-0002",
			SamplingPointID: 1,
			PointNo:         "SP-2024-0001",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			TimePeriodID:    2,
			PeriodName:      "夜间时段(22:00-6:00)",
			MeasurementDate: time.Date(2024, 3, 15, 0, 0, 0, 0, time.Local),
			MeasurementTime: "00:15",
			Leq:             72.3,
			Lmax:            78.9,
			Lmin:            58.1,
			L10:             75.6,
			L90:             60.2,
			StandardLimit:   55.0,
			ExceedValue:     17.3,
			IsExceeded:      true,
			Status:          models.NoiseReadingStatusConfirmed,
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			ReadingNo:       "NR-2024-0003",
			SamplingPointID: 2,
			PointNo:         "SP-2024-0002",
			ComplaintID:     2,
			ComplaintNo:     "CMP-2024-0002",
			TimePeriodID:    1,
			PeriodName:      "昼间时段(6:00-22:00)",
			MeasurementDate: time.Date(2024, 3, 16, 0, 0, 0, 0, time.Local),
			MeasurementTime: "14:30",
			Leq:             62.8,
			Lmax:            68.5,
			Lmin:            55.2,
			L10:             65.3,
			L90:             58.1,
			StandardLimit:   65.0,
			ExceedValue:     -2.2,
			IsExceeded:      false,
			Status:          models.NoiseReadingStatusReviewing,
			ResponsibleUser: "采样员刘明",
			BatchNo:         "BATCH-2024-001",
		},
		{
			ReadingNo:       "NR-2024-0004",
			SamplingPointID: 3,
			PointNo:         "SP-2024-0003",
			ComplaintID:     3,
			ComplaintNo:     "CMP-2024-0003",
			TimePeriodID:    1,
			PeriodName:      "昼间时段(6:00-22:00)",
			MeasurementDate: time.Date(2024, 3, 17, 0, 0, 0, 0, time.Local),
			MeasurementTime: "16:45",
			Leq:             58.2,
			Lmax:            63.7,
			Lmin:            50.8,
			L10:             60.5,
			L90:             53.2,
			StandardLimit:   65.0,
			ExceedValue:     -6.8,
			IsExceeded:      false,
			Status:          models.NoiseReadingStatusConfirmed,
			ResponsibleUser: "采样员张华",
			BatchNo:         "BATCH-2024-002",
		},
		{
			ReadingNo:       "NR-2024-0005",
			SamplingPointID: 5,
			PointNo:         "SP-2024-0005",
			ComplaintID:     5,
			ComplaintNo:     "CMP-2024-0005",
			TimePeriodID:    1,
			PeriodName:      "昼间时段(6:00-22:00)",
			MeasurementDate: time.Date(2024, 3, 18, 0, 0, 0, 0, time.Local),
			MeasurementTime: "12:00",
			Leq:             52.5,
			Lmax:            58.3,
			Lmin:            45.6,
			L10:             55.2,
			L90:             48.1,
			StandardLimit:   65.0,
			ExceedValue:     -12.5,
			IsExceeded:      false,
			Status:          models.NoiseReadingStatusArchived,
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-003",
		},
	}
	db.Create(&readings)
}

func seedEvidenceAttachments(db *gorm.DB) {
	attachments := []models.EvidenceAttachment{
		{
			AttachmentNo:    "EVD-2024-0001",
			FileName:        "noise_measurement_20240315_2330.wav",
			FilePath:        "/uploads/evidence/2024/03/",
			FileSize:        5242880,
			FileType:        "audio/wav",
			MD5Hash:         "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			SamplingPointID: 1,
			PointNo:         "SP-2024-0001",
			NoiseReadingID:  1,
			ReadingNo:       "NR-2024-0001",
			AttachmentType:  "audio_recording",
			Status:          models.EvidenceAttachmentStatusVerified,
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			AttachmentNo:    "EVD-2024-0002",
			FileName:        "site_photo_20240315_2335.jpg",
			FilePath:        "/uploads/evidence/2024/03/",
			FileSize:        2097152,
			FileType:        "image/jpeg",
			MD5Hash:         "b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			SamplingPointID: 1,
			PointNo:         "SP-2024-0001",
			AttachmentType:  "site_photo",
			Status:          models.EvidenceAttachmentStatusVerified,
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			AttachmentNo:    "EVD-2024-0003",
			FileName:        "measurement_report_20240316.pdf",
			FilePath:        "/uploads/evidence/2024/03/",
			FileSize:        1048576,
			FileType:        "application/pdf",
			MD5Hash:         "c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8",
			ComplaintID:     2,
			ComplaintNo:     "CMP-2024-0002",
			SamplingPointID: 2,
			PointNo:         "SP-2024-0002",
			AttachmentType:  "measurement_report",
			Status:          models.EvidenceAttachmentStatusPending,
			ResponsibleUser: "采样员刘明",
			BatchNo:         "BATCH-2024-001",
		},
	}
	db.Create(&attachments)
}

func seedRectificationMeasures(db *gorm.DB) {
	deadline1 := time.Date(2024, 3, 20, 0, 0, 0, 0, time.Local)
	deadline2 := time.Date(2024, 3, 25, 0, 0, 0, 0, time.Local)
	deadline3 := time.Date(2024, 3, 30, 0, 0, 0, 0, time.Local)
	completedAt := time.Date(2024, 3, 28, 0, 0, 0, 0, time.Local)

	measures := []models.RectificationMeasure{
		{
			MeasureNo:       "RM-2024-0001",
			MeasureName:     "调整施工时间避开夜间时段",
			Description:     "将打桩作业时间调整为昼间8:00-18:00，夜间停止施工",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			EnterpriseName:  "华泰建筑有限公司",
			ResponsibleUser: "项目经理张伟",
			Deadline:        &deadline1,
			Status:          models.RectificationMeasureStatusCompleted,
			Effectiveness:   "有效",
			BatchNo:         "BATCH-2024-001",
		},
		{
			MeasureNo:       "RM-2024-0002",
			MeasureName:     "安装风机消声装置",
			Description:     "在风机进出口安装消声器，降低噪声排放",
			ComplaintID:     2,
			ComplaintNo:     "CMP-2024-0002",
			EnterpriseName:  "东方机械制造厂",
			ResponsibleUser: "设备主管李强",
			Deadline:        &deadline2,
			Status:          models.RectificationMeasureStatusImplementing,
			BatchNo:         "BATCH-2024-001",
		},
		{
			MeasureNo:       "RM-2024-0003",
			MeasureName:     "空调外机加装隔音罩",
			Description:     "为空调外机加装隔音罩，减少噪声传播",
			ComplaintID:     3,
			ComplaintNo:     "CMP-2024-0003",
			EnterpriseName:  "万达商业广场",
			ResponsibleUser: "物业经理王芳",
			Deadline:        &deadline3,
			CompletedAt:     &completedAt,
			Status:          models.RectificationMeasureStatusVerified,
			Effectiveness:   "有效",
			BatchNo:         "BATCH-2024-002",
		},
	}
	db.Create(&measures)
}

func seedRetestRecords(db *gorm.DB) {
	records := []models.RetestRecord{
		{
			RetestNo:        "RT-2024-0001",
			ComplaintID:     1,
			ComplaintNo:     "CMP-2024-0001",
			SamplingPointID: 1,
			PointNo:         "SP-2024-0001",
			MeasureID:       1,
			MeasureNo:       "RM-2024-0001",
			RetestDate:      time.Date(2024, 3, 22, 0, 0, 0, 0, time.Local),
			OriginalLeq:     68.5,
			RetestLeq:       52.3,
			ReductionValue:  16.2,
			IsPassed:        true,
			Status:          models.RetestRecordStatusPassed,
			Conclusion:      "整改措施有效，噪声已达标",
			ResponsibleUser: "采样员王强",
			BatchNo:         "BATCH-2024-001",
		},
		{
			RetestNo:        "RT-2024-0002",
			ComplaintID:     3,
			ComplaintNo:     "CMP-2024-0003",
			SamplingPointID: 3,
			PointNo:         "SP-2024-0003",
			MeasureID:       3,
			MeasureNo:       "RM-2024-0003",
			RetestDate:      time.Date(2024, 3, 29, 0, 0, 0, 0, time.Local),
			OriginalLeq:     58.2,
			RetestLeq:       48.5,
			ReductionValue:  9.7,
			IsPassed:        true,
			Status:          models.RetestRecordStatusPassed,
			Conclusion:      "隔音罩安装后噪声明显降低",
			ResponsibleUser: "采样员张华",
			BatchNo:         "BATCH-2024-002",
		},
	}
	db.Create(&records)
}

func seedRuleConfigs(db *gorm.DB) {
	rules := []models.RuleConfig{
		{
			RuleNo:      "RULE-001",
			RuleName:    "夜间噪声超标规则",
			RuleType:    "threshold",
			Description: "夜间时段噪声超过限值10dB以上触发高优先级",
			Conditions:  `{"period_type": "night", "exceed_threshold": 10}`,
			Threshold:   10.0,
			Action:      "set_priority_high",
			Priority:    1,
			Status:      models.RuleConfigStatusActive,
		},
		{
			RuleNo:      "RULE-002",
			RuleName:    "时段不合规规则",
			RuleType:    "compliance",
			Description: "采样时间不在规定时段内触发异常",
			Conditions:  `{"check_time_period": true}`,
			Threshold:   0,
			Action:      "create_anomaly",
			Priority:    2,
			Status:      models.RuleConfigStatusActive,
		},
		{
			RuleNo:      "RULE-003",
			RuleName:    "证据完整度规则",
			RuleType:    "completeness",
			Description: "证据完整度低于80%需要补充",
			Conditions:  `{"min_completeness": 0.8}`,
			Threshold:   0.8,
			Action:      "require_supplement",
			Priority:    3,
			Status:      models.RuleConfigStatusActive,
		},
	}
	db.Create(&rules)
}

func seedAnomalyEvents(db *gorm.DB) {
	deadline := time.Date(2024, 3, 20, 0, 0, 0, 0, time.Local)

	events := []models.AnomalyEvent{
		{
			EventNo:         "ANOM-2024-0001",
			EventName:       "采样时段不合规异常",
			EventType:       "time_period_violation",
			Severity:        models.AnomalyEventSeverityHigh,
			EntityType:      "noise_reading",
			EntityID:        3,
			EntityNo:        "NR-2024-0003",
			TriggerField:    "measurement_time",
			TriggerValue:    "14:30",
			ThresholdValue:  "昼间时段 06:00-22:00",
			Description:     "噪声读数 NR-2024-0003 的测量时间需要复核",
			Status:          models.AnomalyEventStatusOpen,
			ResponsibleUser: "赵处理员",
			Deadline:        &deadline,
			RuleID:          2,
			RuleNo:          "RULE-002",
			BatchNo:         "BATCH-2024-001",
		},
	}
	db.Create(&events)
}
