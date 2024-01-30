# nome_do_app/models.py
from django.db import models
import uuid
from django.utils import timezone

class User(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(default=timezone.now, editable=False)
    updated_at = models.DateTimeField(default=timezone.now)
    name = models.TextField()
    api_key = models.CharField(max_length=64, unique=True, default='')

    class Meta:
        db_table = "User"

    def save(self, *args, **kwargs):
        self.updated_at = timezone.now()
        return super().save(*args, **kwargs)

class Feed(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(default=timezone.now, editable=False)
    updated_at = models.DateTimeField(default=timezone.now)
    name = models.TextField()
    url = models.TextField(unique=True)
    last_fetched_at = models.DateTimeField(null=True, blank=True)

    user = models.ForeignKey('User', on_delete=models.CASCADE)

    class Meta:
        db_table = "Feed"

    def save(self, *args, **kwargs):
        self.updated_at = timezone.now()
        return super().save(*args, **kwargs)
    
class Post(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(default=timezone.now, editable=False)
    updated_at = models.DateTimeField(default=timezone.now)
    title = models.TextField()
    url = models.TextField(unique=True)
    description = models.TextField(blank=True, null=True)
    published_at = models.DateTimeField(blank=True, null=True)
    feed = models.ForeignKey('Feed', on_delete=models.CASCADE)

    class Meta:
        db_table = "Post"

    def save(self, *args, **kwargs):
        self.updated_at = timezone.now()
        return super().save(*args, **kwargs)

